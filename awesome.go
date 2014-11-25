package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type AwesomePlugin struct{}

func (c *AwesomePlugin) Run(cliConnection plugin.CliConnection, args []string) {

	/*
	   	regularlyAwesomeLogo := `
	           ▓▓▓▓▓▓▓   ▓▓▓▓▓▓▓▓▓▓▓
	         ▓▓▓▓▓▓▓▓▓   ▓▓▓▓▓▓▓▓▓▓▓
	        ▓▓▓▓         ▓▓▓▓
	        ▓▓▓▓         ▓▓▓▓▓▓▓▓▓▓▓
	        ▓▓▓▓         ▓▓▓▓▓▓▓▓▓▓▓
	        ▓▓▓▓         ▓▓▓▓
	         ▓▓▓▓▓▓▓▓▓   ▓▓▓▓
	           ▓▓▓▓▓▓▓   ▓▓▓▓
	   	`
	*/

	flamingLogo := `

             )       (       (               )               )  
   (   (  ( /(       )\ )    )\ )         ( /(  (         ( /(  
   )\  )\ )\())   ( (()/(   (()/(      (  )\()) )\ ) (    )\()) 
 (((_)((_|(_)\   ))\ /(_))   /(_))(   ))\((_)\ (()/( )(  ((_)\  
 )\___ _   ((_) /((_|_))_   (_))_|)\ /((_)_((_) ((_)|()\__ ((_) 
((/ __| | / _ \(_))( |   \  | |_ ((_|_))(| \| | _| | ((_) \ / / 
 | (__| || (_) | || || |) | | __/ _ \ || | .\ / _  || ^_|\ V /  
  \___|_| \___/ \_,_||___/  |_| \___/\_,_|_|\_\__,_||_|   |_|   
                                                                
	`

	extraFlamingLogo := `

                                               /
          )         )                         ))        (
    )    /(        ( (         (             /  )       \
   ( \   ) )  /   (   ))       )\       (    )  (      ( (
   ) (  /\ ( )    \ \  ((     ( (       )\  )  \ (     )  \   )  
  (   )(  ( \ (   )  ( )\ )   )\ )     /   ( /(  )(    /  (  /(  
  )\   )\ )\())  (  ( ()/( \ (()/(    (    )\()) )\ ) ( \ )\()) 
 (((_)((_|(_)\   ))\ /(_) )  /(_))(   ) \ (_)\ (()/( )( /((_) \  
 )\___ _   ((_) /((_|_))_)  (_))_|)\ /((_)_((_) ((_)|()\__ ((_) 
((/ __| | / _ \(_))( |   \  | |_ ((_|_))(| \| | _| | ((_) \ / / 
 | (__| || (_) | || || |) | | __/ _ \ || | .\ / _  || ^_|\ V /  
  \___|_| \___/ \_,_||___/  |_| \___/\_,_|_|\_\__,_||_|   |_|   
             
	`

	if args[0] == "awesome" {
		if len(args) >= 2 && args[1] == "--extra-flames" {
			fmt.Println(extraFlamingLogo)
		} else {
			fmt.Println(flamingLogo)
		}
	}

}

func (c *AwesomePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "AwesomePlugin",
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "awesome",
				HelpText: "Awesome? Yes.",
			},
		},
	}
}

func main() {
	plugin.Start(new(AwesomePlugin))
}
