package main

import (
	"fmt"
	"encoding/json"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/hashicorp/terraform/terraform"
)

var datasourceInvalidSchemaElements = map[string]bool{
	"aws_ecs_container_definition":true,
	"aws_vpc_peering_connection":true,
}

var resourceInvalidSchemaElements = map[string]bool{
	"aws_api_gateway_deployment":true,
	"aws_api_gateway_gateway_response":true,
	"aws_api_gateway_integration":true,
	"aws_api_gateway_integration_response":true,
	"aws_api_gateway_method":true,
	"aws_api_gateway_method_response":true,
	"aws_batch_job_definition":true,
	"aws_glue_catalog_database":true,
	"aws_lambda_function":true,
}

func rtNames(rts []terraform.ResourceType) (list []string) {
	for _, rt := range rts {
		if rt.SchemaAvailable && !resourceInvalidSchemaElements[rt.Name] {
			list = append(list, rt.Name)
		}
	}
	return
}
func dsNames(rts []terraform.DataSource) (list []string) {
	for _, rt := range rts {
		if rt.SchemaAvailable && !datasourceInvalidSchemaElements[rt.Name] {
			list = append(list, rt.Name)
		}
	}
	return
}


func main() {
	p := aws.Provider()
	//for _, r  := range p.Resources() {
	//	fmt.Println(r.Name)
	//	if r.SchemaAvailable && !resourceInvalidSchemaElements[r.Name] {
	//		psr := terraform.ProviderSchemaRequest{
	//			ResourceTypes: []string{r.Name},
	//			DataSources: []string{},
	//		}
	//		ps, error := p.GetSchema(&psr)
	//		if error != nil { fmt.Println(error.Error()) }
	//		bolB, _ := json.Marshal(ps)
	//		fmt.Println(string(bolB))
	//	}
	//}
	//for _, r  := range p.DataSources() {
	//	fmt.Println(r.Name)
	//	if r.SchemaAvailable && !datasourceInvalidSchemaElements[r.Name] {
	//		psr := terraform.ProviderSchemaRequest{
	//			ResourceTypes: []string{},
	//			DataSources: []string{r.Name},
	//		}
	//		ps, _ := p.GetSchema(&psr)
	//		bolB, _ := json.Marshal(ps)
	//		fmt.Println(string(bolB))
	//	}
	//}

	psr := terraform.ProviderSchemaRequest{
		ResourceTypes: rtNames(p.Resources()),
		DataSources: dsNames(p.DataSources()),
	}
	ps, _ := p.GetSchema(&psr)

	bolB, _ := json.Marshal(ps)
	fmt.Println(string(bolB))

}