## NetOps For Cisco Devices

## Description
A rest api for configuring and managing multiple cisco devices via a http client json configuration file.

## Badges

## Visuals

## Installation

## Usage Server And Client
```
go run ./cmd/netrasp 

vim configuration.json
{                                                                                                  
    "hostname": ["R1", "R3", "R1"],                                                                
    "username": ["bersen", "bersen", "bersen"],                                                    
    "password": ["bersen", "bersen", "bersen"],                                                    
    "configs": {                                                                                   
        "r1": ["enable", " bersen", "conf t", "end", "sh ip int bri", "disable"],                  
        "r3": ["enable", " bersen", "conf t", "end", "sh ip arp", "disable"],                      
        "r11": ["enable", " bersen", "conf t", "end", "sh ip int bri", "disable"]                  
    }                                                                                              
}      

go run ./cmd/client/ -url http://localhost:3000/create/config -filename configuration.json
```

## Support

## Roadmap

## Contributing

## Authors and acknowledgment
Bersen Naidoo

## License

## Project status
WIP

