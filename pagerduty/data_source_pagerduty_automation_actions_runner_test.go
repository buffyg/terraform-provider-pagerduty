package pagerduty

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourcePagerDutyAutomationActionsRunner_Basic(t *testing.T) {
	// name := fmt.Sprintf("tf-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePagerDutyAutomationActionsRunnerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPagerdutyAutomationActionsRunnerExists("data.pagerduty_automation_actions_runner.foo"),
				),
			},
		},
	})
}

func testAccCheckPagerdutyAutomationActionsRunnerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// client, _ := testAccProvider.Meta().(*Config).Client()
		// client.AutomationActionsRunner.Create()

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Runner ID is set")
		}

		return nil
	}
}

func testAccDataSourcePagerDutyAutomationActionsRunnerConfig() string {
	return fmt.Sprintf(``)
}
