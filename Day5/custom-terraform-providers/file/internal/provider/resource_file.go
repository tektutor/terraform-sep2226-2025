package provider

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocalFile() *schema.Resource {
	return &schema.Resource{
		Description: "This resource will support Create, Read, Update and Delete file resource via Terraform.",

		CreateContext: resourceCreateFile,
		ReadContext:   resourceReadFile,
		UpdateContext: resourceUpdateFile,
		DeleteContext: resourceDeleteFile,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Description: "Name of the file",
				Type:        schema.TypeString,
				Required:    true,
			},

			"file_content": {
				Description: "Content that must be stored/retrieved to/from the file.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	//Retrieve the input user provided in the terraform .tf file
	filename := d.Get("file_name").(string)
	filecontent := d.Get("file_content").(string)

	myFile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	myFile.WriteString(filecontent)
	myFile.Sync() //This will ensure the file content is flushed out to disk immediately

	d.SetId("File123")

	return nil
}

func resourceReadFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve the input user provided in the terraform .tf file
	filename := d.Get("file_name").(string)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	d.Set("content", content)

	return nil
}

func resourceUpdateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve the input user provided in the terraform .tf file
	filename := d.Get("file_name").(string)
	content := d.Get("file_content").(string)
	id := d.Id()

	myFile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	myFile.WriteString(content + "\n")
	myFile.Sync() //This will ensure the file content is flushed out to disk immediately
	d.SetId(id)

	return nil
}

func resourceDeleteFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve the input user provided in the terraform .tf file
	filename := d.Get("file_name").(string)

	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}

	return nil
}
