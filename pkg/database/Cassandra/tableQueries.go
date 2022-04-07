package cassandra

var tableQueries = [1]string{
	`CREATE TABLE IF NOT EXISTS ProjectMetadata.customer_projects(id bigint, project_id bigint, industry_id smallint, product_id smallint, created_at timestamp, status boolean, PRIMARY KEY ((project_id, industry_id, product_id), created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
}

func GetTableQueries() [1]string {
	return tableQueries
}
