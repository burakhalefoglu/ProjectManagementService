package Cassandra

import (
	"ProjectManagementService/internal/model"
	cassandra "ProjectManagementService/pkg/database/Cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassCustomerProjectDal struct {
	Client *gocql.Session
	Table  string
}

func NewCustomerProjectDal(Table string) *cassCustomerProjectDal {
	return &cassCustomerProjectDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassCustomerProjectDal) Add(data *model.CustomerProject) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO "+
		"%s(id, project_id, industry_id, product_id, created_at, status)"+
		" VALUES(?,?,?,?,?,?)", m.Table),
		data.Id, data.ProjectId, data.IndustryId, data.ProductId, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}

func (m *cassCustomerProjectDal) UpdateByPK(data *model.CustomerProject) error {

	if err := m.Client.Query(fmt.Sprintf("UPDATE %s SET"+
		" id = ?, project_id = ?, industry_id = ?  product_id = ? created_at = ? status = ?"+
		" WHERE project_id = %d AND industry_id = %d AND product_id = %d", m.Table, data.ProjectId, data.IndustryId, data.ProductId),
		data.Id, data.ProjectId, data.IndustryId, data.ProductId, data.CreatedAt, data.Status).Exec(); err != nil {
		return err
	}
	return nil
}

func (m *cassCustomerProjectDal) GetByPK(projectId int64, industryId int8, productId int8) (*model.CustomerProject, error) {

	data := &model.CustomerProject{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE project_id = %d AND industry_id = %d AND product_id = %d LIMIT 1",
		m.Table, projectId, industryId, productId)).
		Scan(&data.Id, &data.ProjectId, &data.IndustryId, &data.ProductId, &data.CreatedAt, &data.Status); err != nil {
		return nil, err
	}
	return data, nil
}
