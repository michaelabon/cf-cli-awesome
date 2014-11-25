package main

import (
	"fmt"
	"os"

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
         )\___ _   ((_) /((_|_))_/  (_))_|)\ /((_)_((_) ((_)|()\__ ((_)
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
	nfoText := `
 /============================================================================\
 |                     -*- Cloud Foundry is AWESOME -*-                       |
 |                                                                            |  
 |                          (c) The CF Foundation                             |
 ├----------------------------------------------------------------------------┤
 |  Date       :  November 25th, 2014                                         |
 |  Cracker    :  The Kenyonator              Control :  Keyboard             |
 |  Supplier   :  Branch Davidian             Display :  Hercules/CGA/EGA/VGA |
 |  Protection :  None                        Sound   :  Adlib/SoundBlaster   |
 |                                                                            |
 \----------------------------------------------------------------------------/
 /----------------------------------------------------------------------------\
 |                                                                            |
 |  Personal ones to : The Mighty Morhovich, Shurst                           |
 |                                                                            |
 |  Group HiHo's to : The CLIsters                                            |
 |                                                                            |
 \----------------------------------------------------------------------------/
 /----------------------------------------------------------------------------\
 |                                                                            |
 |   Give a buzz to our wellknown Bulletin Board Systems around the world...  |
 |                                                                            |
 |      WE ARE CURRENTLY LOOKING FOR MEMBER BOARDS AND DISTRIBUTION SITES     |   
 |                                                                            |
 |                                                                            |
 |                 Call      - 1-646-792-5770                                 |
 |                 Log In As - SQUIRREL                                       |
 |                 Password  - MAKEITHAPPEN                                   |                           
 |                                                                            |
 |                                                                            |
 |                                                                            |
 |            Interested in Hadoop Clusters?                                  |
 |            For the best prices contact    1-650-846-1600                   |
 |                                                                            |
 |                                                                            |
 ├----------------------------------------------------------------------------┤
 |                          ThERE CAN BE oNLY oNE...                          |
 \----------------------------------------------------------------------------/

        SUPPORT THE COMPANIES THAT PROGRAM AND CODE QUALITY SOFTWARE (c).
                     IF YOU LIKED THIS CLOUD FOUNDRY, BUY IT.
                               -*- PC/NYCF -*-
    `

	if args[0] == "awesome" {
		extraFlames := false
		nfo := false
		for _, arg := range args {
			switch arg {
			case "--extra-flames":
				extraFlames = true
			case "--nfo":
				nfo = true
			case "awesome":
			default:
				fmt.Println("Your usage is not awesome enough.  Try --extra-flames or --nfo.")
				os.Exit(1)
			}
		}

		if extraFlames {
			fmt.Println(extraFlamingLogo)
		} else {
			fmt.Println(flamingLogo)
		}

		if nfo {
			fmt.Println(nfoText)
		}
	}

}

func (c *AwesomePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "AwesomePlugin",
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "awesome",
				HelpText: "Awesome? Yes. Use --extra-flames when needed.",
			},
		},
	}
}

func main() {
	plugin.Start(new(AwesomePlugin))
}
