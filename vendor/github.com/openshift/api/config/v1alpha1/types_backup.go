package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
// Backup provides configuration for performing backups of the openshift cluster.
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type Backup struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
	// +required
	Spec BackupSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +kubebuilder:validation:Optional
	// +optional
	Status BackupStatus `json:"status"`
}

type BackupSpec struct {
	// etcd specifies the configuration for periodic backups of the etcd cluster
	// +kubebuilder:validation:Required
	EtcdBackupSpec EtcdBackupSpec `json:"etcd"`
}

type BackupStatus struct {
}

// EtcdBackupSpec provides configuration for automated etcd backups to the cluster-etcd-operator
type EtcdBackupSpec struct {

	// Schedule defines the recurring backup schedule in Cron format
	// every 2 hours: 0 */2 * * *
	// every day at 3am: 0 3 * * *
	// Setting to an empty string "" means disabling scheduled backups
	// +kubebuilder:validation:Optional
	// +optional
	// +kubebuilder:validation:Pattern:=`^(@(annually|yearly|monthly|weekly|daily|hourly))|(\*|(?:\*|(?:[0-9]|(?:[1-5][0-9])))\/(?:[0-9]|(?:[1-5][0-9]))|(?:[0-9]|(?:[1-5][0-9]))(?:(?:\-[0-9]|\-(?:[1-5][0-9]))?|(?:\,(?:[0-9]|(?:[1-5][0-9])))*)) (\*|(?:\*|(?:\*|(?:[0-9]|1[0-9]|2[0-3])))\/(?:[0-9]|1[0-9]|2[0-3])|(?:[0-9]|1[0-9]|2[0-3])(?:(?:\-(?:[0-9]|1[0-9]|2[0-3]))?|(?:\,(?:[0-9]|1[0-9]|2[0-3]))*)) (\*|(?:[1-9]|(?:[12][0-9])|3[01])(?:(?:\-(?:[1-9]|(?:[12][0-9])|3[01]))?|(?:\,(?:[1-9]|(?:[12][0-9])|3[01]))*)) (\*|(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)(?:(?:\-(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC))?|(?:\,(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC))*)) (\*|(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT)(?:(?:\-(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT))?|(?:\,(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT))*))$`
	Schedule string `json:"schedule"`

	// Cron Regex breakdown:
	// Allow macros: (@(annually|yearly|monthly|weekly|daily|hourly))
	// OR
	// Minute:
	//   (\*|(?:\*|(?:[0-9]|(?:[1-5][0-9])))\/(?:[0-9]|(?:[1-5][0-9]))|(?:[0-9]|(?:[1-5][0-9]))(?:(?:\-[0-9]|\-(?:[1-5][0-9]))?|(?:\,(?:[0-9]|(?:[1-5][0-9])))*))
	// Hour:
	//   (\*|(?:\*|(?:\*|(?:[0-9]|1[0-9]|2[0-3])))\/(?:[0-9]|1[0-9]|2[0-3])|(?:[0-9]|1[0-9]|2[0-3])(?:(?:\-(?:[0-9]|1[0-9]|2[0-3]))?|(?:\,(?:[0-9]|1[0-9]|2[0-3]))*))
	// Day of the Month:
	//   (\*|(?:[1-9]|(?:[12][0-9])|3[01])(?:(?:\-(?:[1-9]|(?:[12][0-9])|3[01]))?|(?:\,(?:[1-9]|(?:[12][0-9])|3[01]))*))
	// Month:
	//   (\*|(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)(?:(?:\-(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC))?|(?:\,(?:[1-9]|1[012]|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC))*))
	// Day of Week:
	//   (\*|(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT)(?:(?:\-(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT))?|(?:\,(?:[0-6]|SUN|MON|TUE|WED|THU|FRI|SAT))*))
	//

	// The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	// If not specified, this will default to the time zone of the kube-controller-manager process.
	// See https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones
	// +kubebuilder:validation:Optional
	// +optional
	// +kubebuilder:validation:Pattern:=`^([A-Za-z_]+([+-]*0)*|[A-Za-z_]+(\/[A-Za-z_]+){1,2})(\/GMT[+-]\d{1,2})?$`
	TimeZone string `json:"timeZone,omitempty"`

	// Timezone regex breakdown:
	// ([A-Za-z_]+([+-]*0)*|[A-Za-z_]+(/[A-Za-z_]+){1,2}) - Matches either:
	//   [A-Za-z_]+([+-]*0)* - One or more alphabetical characters (uppercase or lowercase) or underscores, followed by a +0 or -0 to account for GMT+0 or GMT-0 (for the first part of the timezone identifier).
	//   [A-Za-z_]+(/[A-Za-z_]+){1,2} - One or more alphabetical characters (uppercase or lowercase) or underscores, followed by one or two occurrences of a forward slash followed by one or more alphabetical characters or underscores. This allows for matching timezone identifiers with 2 or 3 parts, e.g America/Argentina/Buenos_Aires
	// (/GMT[+-]\d{1,2})? - Makes the GMT offset suffix optional. It matches "/GMT" followed by either a plus ("+") or minus ("-") sign and one or two digits (the GMT offset)

	// RetentionPolicy defines the retention policy for retaining and deleting existing backups.
	// +kubebuilder:validation:Optional
	// +optional
	RetentionPolicy RetentionPolicy `json:"retentionPolicy"`

	// PVCName specifies the name of the PersistentVolumeClaim which binds a PersistentVolume where the
	// etcd backup files would be saved
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +required
	PVCName string `json:"pvcName"`
}

// RetentionType is the enumeration of valid retention policy types
// +enum
// +kubebuilder:validation:Enum:="RetentionNumber";"RetentionSize"
type RetentionType string

const (
	// RetentionTypeNumber sets the retention policy based on the number of backup files saved
	RetentionTypeNumber RetentionType = "RetentionNumber"
	// RetentionTypeSize sets the retention policy based on the total size of the backup files saved
	RetentionTypeSize RetentionType = "RetentionSize"
)

// RetentionPolicy defines the retention policy for retaining and deleting existing backups.
// This struct is a discriminated union that allows users to select the type of retention policy from the supported types.
// +union
type RetentionPolicy struct {
	// RetentionType sets the type of retention policy.
	// Currently, the only valid policies are retention by number of backups (RetentionNumber), by the size of backups (RetentionSize). More policies or types may be added in the future.
	// Specifying an empty string "" means no retention policy and no backups will be pruned.
	// +unionDiscriminator
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:="";"RetentionNumber";"RetentionSize"
	RetentionType RetentionType `json:"retentionType"`

	// RetentionNumber configures the retention policy based on the number of backups
	// +kubebuilder:validation:Optional
	// +optional
	RetentionNumber *RetentionNumberConfig `json:"retentionNumber,omitempty"`

	// RetentionSize configures the retention policy based on the size of backups
	// +kubebuilder:validation:Optional
	// +optional
	RetentionSize *RetentionSizeConfig `json:"retentionSize,omitempty"`
}

// RetentionNumberConfig specifies the configuration of the retention policy on the number of backups
type RetentionNumberConfig struct {
	// MaxNumberOfBackups defines the maximum number of backups to retain.
	// If the existing number of backups saved is equal to MaxNumberOfBackups then
	// the oldest backup will be removed before a new backup is initiated.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Required
	// +required
	MaxNumberOfBackups int `json:"maxNumberOfBackups,omitempty"`
}

// RetentionSizeConfig specifies the configuration of the retention policy on the total size of backups
type RetentionSizeConfig struct {
	// MaxSizeOfBackupsGb defines the total size in GB of backups to retain.
	// If the current total size backups exceeds MaxSizeOfBackupsGb then
	// the oldest backup will be removed before a new backup is initiated.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Required
	// +required
	MaxSizeOfBackupsGb int `json:"maxSizeOfBackupsGb,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupList is a collection of items
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type BackupList struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard list's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`
	Items           []Backup `json:"items"`
}
