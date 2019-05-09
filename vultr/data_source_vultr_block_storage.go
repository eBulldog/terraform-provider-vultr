package vultr

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vultr/govultr"
)

func dataSourceVultrBlockStorage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVultrBlockStorageRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_pre_month": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"attached_to_vps": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"label": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVultrBlockStorageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client).govultrClient()

	filters, filtersOk := d.GetOk("filter")

	if !filtersOk {
		return fmt.Errorf("issue with filter: %v", filtersOk)
	}

	block, err := client.BlockStorage.GetList(context.Background())

	if err != nil {
		return fmt.Errorf("error getting applications: %v", err)
	}

	blockList := []govultr.BlockStorage{}

	f := buildVultrDataSourceFilter(filters.(*schema.Set))

	for _, b := range block {
		sm, err := structToMap(b)

		if err != nil {
			return err
		}

		if filterLoop(f, sm) {
			blockList = append(blockList, b)
		}
	}

	if len(blockList) > 1 {
		return errors.New("your search returned too many results. Please refine your search to be more specific")
	}

	if len(blockList) < 1 {
		return errors.New("no results were found")
	}

	d.SetId(blockList[0].BlockStorageID)
	d.Set("date_created", blockList[0].DateCreated)
	d.Set("cost_pre_month", blockList[0].Cost)
	d.Set("status", blockList[0].Status)
	d.Set("size_gb", blockList[0].Size)
	d.Set("region_id", blockList[0].RegionID)
	d.Set("attached_to_vps", blockList[0].VpsID)
	d.Set("label", blockList[0].Label)
	return nil
}
