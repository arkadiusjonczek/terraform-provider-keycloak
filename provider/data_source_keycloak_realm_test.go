package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccKeycloakDataSourceRealm_basic(t *testing.T) {
	realm := acctest.RandomWithPrefix("tf-acc")

	resourceName := "keycloak_realm.my_realm"
	dataSourceName := "data.keycloak_realm.realm"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		CheckDestroy:      testAccCheckKeycloakRealmDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceKeycloakRealm_basic(realm),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "realm", resourceName, "realm"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name_html", resourceName, "display_name_html"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		CheckDestroy:      testAccCheckKeycloakRealmDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceKeycloakRealm_basic2(realm),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "realm", resourceName, "realm"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name_html", resourceName, "display_name_html"),
					resource.TestCheckResourceAttrPair(dataSourceName, "default_roles", resourceName, "default_roles"),
				),
			},
		},
	})
}

func testDataSourceKeycloakRealm_basic(realm string) string {
	return fmt.Sprintf(`
resource "keycloak_realm" "my_realm" {
	realm             = "%s"
	enabled           = true
	display_name      = "foo"
	display_name_html = "<b>foo</b>"
}

resource "keycloak_role" "testrole" {
    realm_id = resource.keycloak_realm.my_realm.id
    name     = "testrole"
}

data "keycloak_realm" "realm" {
	realm = keycloak_realm.my_realm.id
}`, realm)
}

func testDataSourceKeycloakRealm_basic2(realm string) string {
	return fmt.Sprintf(`
resource "keycloak_realm" "my_realm" {
	realm             = "%s"
	enabled           = true
	display_name      = "foo"
	display_name_html = "<b>foo</b>"
	default_roles     = ["testrole"]
}

data "keycloak_realm" "realm" {
	realm = keycloak_realm.my_realm.id
}`, realm)
}
