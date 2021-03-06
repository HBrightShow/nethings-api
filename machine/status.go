package machine

import (
	"strings"
	"fmt"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	

	"github.com/nethings/internal-api/machine/types"
	"github.com/nethings/internal-api/query"
	"github.com/nethings/internal-api/response"
)


func (g *Machine) GetStatus(ctx *gin.Context) {

	//fmt.Println("GetStatus recv")
	//ctx.String(200, "Hello, Geektutu")


	builder := query.New()

	support := strings.ToLower(ctx.DefaultQuery("support", ""))
	if support != "" {
		if support != "true" && support != "false" {
			response.BadRequest(ctx, errors.New("wrong value for support parameter"))
			return
		}
		builder.Filters.Add("support", support)
	}


	q, params := builder.Run(`
		SELECT * FROM address
	`)

	q = strings.Replace(q, "$param_overwrite$", fmt.Sprintf("$%d", len(params)), 1)

	//var count int
	rows, err := g.db.Connection().Query(q, params...)
	//fmt.Println("rows: ", rows)

	if err != nil {
		response.Error(ctx, err)
		return
	}
	defer rows.Close()

	
	//fmt.Printf("%T", rows)	//sql.rows

	var albums []types.Album

	// Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb types.Album
        if err := rows.Scan(&alb.Id, &alb.Uid, &alb.Phone,
            &alb.Name, &alb.Zipcode, &alb.Address, &alb.Default_address, &alb.Add_time); err != nil {
            return 
        }
        albums = append(albums, alb)
    }
    if err = rows.Err(); err != nil {
        return 
    }

	//response.OK(ctx, "bright")
	response.OK(ctx, albums)
}