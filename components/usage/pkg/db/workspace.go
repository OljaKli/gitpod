// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package db

import (
	"database/sql"
	"time"
)

// Workspace struct is a row record of the d_b_workspace table in the gitpod database
type Workspace struct {
	//[ 0] id                                             char(36)             null: false  primary: true   isArray: false  auto: false  col: char            len: 36      default: []
	ID string `gorm:"primary_key;column:id;type:char;size:36;" json:"id"`
	//[ 1] creationTime                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CreationTime string `gorm:"column:creationTime;type:varchar;size:255;" json:"creation_time"`
	//[ 2] ownerId                                        char(36)             null: false  primary: false  isArray: false  auto: false  col: char            len: 36      default: []
	OwnerID string `gorm:"column:ownerId;type:char;size:36;" json:"owner_id"`
	//[ 3] contextURL                                     text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ContextURL string `gorm:"column:contextURL;type:text;size:65535;" json:"context_url"`
	//[ 4] description                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description string `gorm:"column:description;type:varchar;size:255;" json:"description"`
	//[ 5] context                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Context string `gorm:"column:context;type:text;size:65535;" json:"context"`
	//[ 6] config                                         text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Config string `gorm:"column:config;type:text;size:65535;" json:"config"`
	//[ 7] imageSource                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ImageSource sql.NullString `gorm:"column:imageSource;type:text;size:65535;" json:"image_source"`
	//[ 8] imageNameResolved                              varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ImageNameResolved string `gorm:"column:imageNameResolved;type:varchar;size:255;" json:"image_name_resolved"`
	//[ 9] archived                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Archived int32 `gorm:"column:archived;type:tinyint;default:0;" json:"archived"`
	//[10] shareable                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Shareable int32 `gorm:"column:shareable;type:tinyint;default:0;" json:"shareable"`
	//[11] _lastModified                                  timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP(6)]
	LastModified time.Time `gorm:"column:_lastModified;type:timestamp;default:CURRENT_TIMESTAMP(6);" json:"_last_modified"`
	//[12] deleted                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Deleted int32 `gorm:"column:deleted;type:tinyint;default:0;" json:"deleted"`
	//[13] type                                           char(16)             null: false  primary: false  isArray: false  auto: false  col: char            len: 16      default: [regular]
	Type string `gorm:"column:type;type:char;size:16;default:regular;" json:"type"`
	//[14] baseImageNameResolved                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BaseImageNameResolved string `gorm:"column:baseImageNameResolved;type:varchar;size:255;" json:"base_image_name_resolved"`
	//[15] softDeleted                                    char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	SoftDeleted sql.NullString `gorm:"column:softDeleted;type:char;size:4;" json:"soft_deleted"`
	//[16] pinned                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Pinned int32 `gorm:"column:pinned;type:tinyint;default:0;" json:"pinned"`
	//[17] softDeletedTime                                varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SoftDeletedTime string `gorm:"column:softDeletedTime;type:varchar;size:255;" json:"soft_deleted_time"`
	//[18] contentDeletedTime                             varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContentDeletedTime string `gorm:"column:contentDeletedTime;type:varchar;size:255;" json:"content_deleted_time"`
	//[19] basedOnPrebuildId                              char(36)             null: true   primary: false  isArray: false  auto: false  col: char            len: 36      default: []
	BasedOnPrebuildID sql.NullString `gorm:"column:basedOnPrebuildId;type:char;size:36;" json:"based_on_prebuild_id"`
	//[20] basedOnSnapshotId                              char(36)             null: true   primary: false  isArray: false  auto: false  col: char            len: 36      default: []
	BasedOnSnapshotID sql.NullString `gorm:"column:basedOnSnapshotId;type:char;size:36;" json:"based_on_snapshot_id"`
	//[21] projectId                                      char(36)             null: true   primary: false  isArray: false  auto: false  col: char            len: 36      default: []
	ProjectID sql.NullString `gorm:"column:projectId;type:char;size:36;" json:"project_id"`
	//[22] cloneURL                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CloneURL string `gorm:"column:cloneURL;type:varchar;size:255;" json:"clone_url"`
}

// TableName sets the insert table name for this struct type
func (d *Workspace) TableName() string {
	return "d_b_workspace"
}
