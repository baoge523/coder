## solidity

### visibility and getters

#### state variable visibility
public

internal (this is the default visibility level for state variable)
> current contracts and derived contracts

private
> private state variable are likes internal ons but they are not visible in derived contracts 


#### function visibility
external (not access internal function)
public  (all)
internal (current contract and derived contracts from it)
private (only current contract)