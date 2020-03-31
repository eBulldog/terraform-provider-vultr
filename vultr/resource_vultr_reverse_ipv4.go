package vultr

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vultr/govultr"
)

func resourceVultrReverseIPV4() *schema.Resource {
	return &schema.Resource{
		Create: resourceVultrReverseIPV4Create,
		Read:   resourceVultrReverseIPV4Read,
		Delete: resourceVultrReverseIPV4Delete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"reverse": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVultrReverseIPV4Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client).govultrClient()

	instanceID := d.Get("instance_id").(string)

	ip, err := getCanonicalIPV4(d.Get("ip").(string))
	if err != nil {
		return err
	}

	reverse := d.Get("reverse").(string)

	log.Printf("[INFO] Creating reverse IPv4")
	err = client.Server.SetReverseIPV4(context.Background(), instanceID, ip, reverse)

	if err != nil {
		return fmt.Errorf("Error creating reverse IPv4: %v", err)
	}

	d.SetId(ip)

	return resourceVultrReverseIPV4Read(d, meta)
}

func resourceVultrReverseIPV4Read(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*Client).govultrClient()

        instanceID := d.Get("instance_id").(string)

        reverseIPV4s, err := client.Server.IPV4Info(context.Background(), instanceID, false)
        if err != nil {
		return fmt.Errorf("Error getting reverse IPv4Info: %v", err)
        }

        var reverseIPV4 *govultr.IPV4
        for i := range reverseIPV4s {
                if reverseIPV4s[i].IP == d.Id() {
                        reverseIPV4 = &reverseIPV4s[i]
                        break
                }
        }

        d.Set("ip", reverseIPV4.IP)
        d.Set("reverse", reverseIPV4.Reverse)

	return nil
}

func resourceVultrReverseIPV4Delete(d *schema.ResourceData, meta interface{}) error {
        client := meta.(*Client).govultrClient()

        instanceID := d.Get("instance_id").(string)

        ip, err := getCanonicalIPV4(d.Get("ip").(string))
        if err != nil {
                return err
        }

        log.Printf("[INFO] Resetting to default reverse IPv4: %s", d.Id())
        err2 := client.Server.SetDefaultReverseIPV4(context.Background(), instanceID, ip)

        if err2 != nil {
                return fmt.Errorf("Error resetting reverse IPv4 (%s): %v", d.Id(), err)
        }

        return nil
}

func getCanonicalIPV4(value string) (string, error) {
	ip := net.ParseIP(value)
	if ip == nil {
		return "", fmt.Errorf("Invalid IPv4 address: %s", value)
	}

	return ip.String(), nil
}
