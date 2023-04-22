Multiple variables of the same type can also be declared on a single line: var x, y int makes x and y both int variables. 
You can also make use of parallel assignment: a, b := 20, 16 If you are using an initializer expression for declaring 
variables, you can omit the type using short variable declaration as shown here:

country, state := "Germany", "Berlin"

We use the operator : = for declaring and initializing variables with short variable declaration. When you declare variables with this method, 
you can't specify the type because the type is determined by the initializer expression.

S