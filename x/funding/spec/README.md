# Funding Module Specification

## Abstract

This module provides the functionality of allowing anyone to buy FUELs using ATOMs using a bonding curve, similar to Bancor Protocol's.

The module makes sure that the following interactions can be achieved.

1. A token (FUEL) can be minted (bought) at any time according to a price set by the bonding curve.
2. This price increases as token supply grows.
3. The money (ATOM) paid for tokens is kept in a reserve pool.
4. At any point in time, a token can be burned (sold) back to the contract, the the money can be redeemed from the reserve pool.


## Contents

// TODO: Create the below files if they are needed.
1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
4. **[Begin-Block](04_begin_block.md)**
5. **[End-Block](06_end_bloc.md)**
6. **[05_hooks](06_hooks.md)**
7. **[Events](07_events.md)**
8. **[Parameters](08_params.md)**
