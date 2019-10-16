# ![ACT logo](../../Documentation/images/ACT_logo_50px.png) Automatic Component Toolkit

## Example: Cross Component Inheritance

### act Numbers.xml
Creates Numbers_component which simply shows the "new" ABI that is based on Extended Handles.


Numbers_component already implements a generated Numbers component.

### act Calculation.xml
Creates Calculation_component which imports the Numbers component and shows how the new ABI of Numbers is passed through the ABI of Calculation.

Calculation_component already implements a generated Calculation component.

## TODO:

### act SpecialNumbers.xml
Creates SpecialNumbers_component with a class `SpecialVariable` that inherits from the `Variable` Class in the Numbers component.

Once this is implemented, I will try to feed the Calculation-component an instance of `SpecialVariable`.








