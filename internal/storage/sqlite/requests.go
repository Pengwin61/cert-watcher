package sqlite

// func (c *Client) Insert(notBefore string, notAfter, organization, OU, commonName string) {

// 	_, err := c.Db.Exec("insert into certificate (before, after, organization, ou, cn) values ($1, $2, $3, $4, $5)",
// 		notBefore, notAfter, organization, OU, commonName)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (c *Client) Get() {
// 	rows, err := c.Db.Query("select * from certificate")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	certificates := []certificate{}

// 	for rows.Next() {
// 		c := certificate{}
// 		err := rows.Scan(&c.id, &c.before, &c.after, &c.organization, &c.ou, &c.cn)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		certificates = append(certificates, c)
// 	}
// 	for _, c := range certificates {

// 		fmt.Println(c.id, c.before, c.after, c.organization, c.ou, c.cn)
// 	}
// }

// func (c *Client) Update(id int) {

// 	_, err := c.Db.Exec("update certificate set organization = $1 where id = $2", "", id)
// 	if err != nil {
// 		panic(err)
// 	}
// }
