package vthunder

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_SLB_TEMPLATE_MONITOR_RESOURCE = `
resource "vthunder_slb_template_monitor" "testname" {
	user_tag = "test_tag"
	link_disable_cfg {
		dis_sequence = 1 
		diseth = 1
		}
	monitor_relation = "monitor-and"
	id2 = 1
}
`

//Acceptance test
func TestSlbTemplateMonitor_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TEST_SLB_TEMPLATE_MONITOR_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "name", "testmonitor"),
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "user_tag", "test_tag"),
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "link_disable_cfg.0.dis_sequence", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "link_disable_cfg.0.diseth", "1"),
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "monitor_relation", "monitor-and"),
					resource.TestCheckResourceAttr("vthunder_slb_template_monitor.testname", "id", "1"),
				),
			},
		},
	})
}
