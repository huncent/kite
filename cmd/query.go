package cmd

import (
	"flag"
	"fmt"
	"kite"
	"kite/cmd/util"
	"kite/protocol"
	"net/url"
)

type Query struct {
	client *kite.Kite
}

func NewQuery(client *kite.Kite) *Query {
	return &Query{
		client: client,
	}
}

func (r *Query) Definition() string {
	return "Query kontrol"
}

func (r *Query) Exec(args []string) error {
	token, err := util.ParseKiteKey()
	if err != nil {
		return err
	}

	username, _ := token.Claims["sub"].(string)

	var query protocol.KontrolQuery
	flags := flag.NewFlagSet("query", flag.ContinueOnError)
	flags.StringVar(&query.Username, "username", username, "")
	flags.StringVar(&query.Environment, "environment", "", "")
	flags.StringVar(&query.Name, "name", "", "")
	flags.StringVar(&query.Version, "version", "", "")
	flags.StringVar(&query.Region, "region", "", "")
	flags.StringVar(&query.Hostname, "hostname", "", "")
	flags.StringVar(&query.ID, "id", "", "")
	flags.Parse(args)

	parsed, err := url.Parse(token.Claims["kontrolURL"].(string))
	if err != nil {
		return err
	}

	kontrol := r.client.NewKontrol(parsed)
	if err = kontrol.Dial(); err != nil {
		return err
	}

	response, err := kontrol.Tell("getKites", query)
	if err != nil {
		return err
	}

	var kites []protocol.KiteWithToken
	err = response.Unmarshal(&kites)
	if err != nil {
		return err
	}

	for i, kite := range kites {
		fmt.Printf("\t%d.\t%+v\n", i+1, kite.Kite)
	}

	return nil
}