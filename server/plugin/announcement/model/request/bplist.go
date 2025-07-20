
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type BplistSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ScheduleId  *string `json:"scheduleId" form:"scheduleId"` 
       TeamId  *string `json:"teamId" form:"teamId"` 
    request.PageInfo
}
