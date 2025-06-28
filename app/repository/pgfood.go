package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetFoodRow struct {
	ID           int
	Name         string
	Type         pgtype.Text
	IntakeStatus pgtype.Text
	Feeder       pgtype.Text
	Location     pgtype.Text
	Remarks      pgtype.Text
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

func (p *pgRepoImpl) GetFood(id int) (GetFoodRow, error) {
	query := `
select
    f.id,
    f.name,
    ft.name as type,
    fis.name as intake_status,
    ff.name as feeder,
    fl.name as location,
    f.remarks,
    f.created_at,
    f.updated_at
from food f
left join food_types ft on f.type_id = ft.id
left join food_intake_status fis on f.intake_status_id = fis.id
left join food_feeders ff on f.feeder_id = ff.id
left join food_locations fl on f.location_id = fl.id
where f.id = $1
and f.removed_at is null
	`
	rows := p.pgxPool.QueryRow(context.Background(), query, id)
	result := GetFoodRow{}
	err := rows.Scan(
		&result.ID,
		&result.Name,
		&result.Type,
		&result.IntakeStatus,
		&result.Feeder,
		&result.Location,
		&result.Remarks,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	return result, err
}
