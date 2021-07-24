# Watering System Notes

## Pump pinout

By default the app is configured to expect to able to control the pump by setting GPIO4 to High (1). If unsure, on RaspberryPi devices you can find it using the pinout command. More information is available in [RaspberryPi GPIO Documentation](https://www.raspberrypi.org/documentation/usage/gpio/)

## Materials

- **Water tank/canister**
    - Example: [60l canister](https://www.plasteelaste.de/60-liter-kanister-mit-3-griffen-in-naturweiss-oder-blau)
    - Remarks: Make sure it is sized for your needs. Size, volume, etc.

- **Pump**
    - Example: [Javtop JT-750 DC12V](https://www.aliexpress.com/item/33030972301.html)
    - Remarks: Make sure the pump:
        - has sufficient power - maybe at least double max waterlevel (H-Max) compared to your needs.
        - fits in the water tank.
        - comes with sufficient cable lenght to avoid connections in wet areas.
        - cosnider buying pump with already installed connector.

- **Water hoses and connectors**
    - Example: [Multi-Purpose Hose Adaptor, Laguna](https://www.amazon.de/-/en/gp/product/B001BO689Q/) & 8mm water hose

    - Remarks: pump may have different diameter to your canister and main water hose. take in account and plan adaptor, if needed.

- **Reley**
    - Example: [JQC-3FF-S-Z](https://www.aliexpress.com/item/4001237707276.html)
    - Remarks: Make sure the relay
        - has a jumper for defining the trigger level. Has to be set to high to work with the code of the project.
        - has the right operating voltage and maximum load for your load.

- **Power Adapter**
    - Example: [CLS-832](https://www.aliexpress.com/item/1005001982270053.html) or [MANUAL-30W + USB](https://www.aliexpress.com/item/4000046315142.html)
    - Remarks:
        - Needs to be able to power your pump. Charging the pump and the raspberry or other embedded board from same cable may be problematic. Consider adapter with separate usb outlate for powering the raspberry.

- **Case and external power cord**
    - Example: [Extension Cable with Euro Plug](https://www.amazon.de/-/en/P01303-Extension-Coupling-Plastic-Indoor/dp/B07Y8Q2F5D)
    - Remarks: 
        - Consider the power consumption of the pump. You may need 
