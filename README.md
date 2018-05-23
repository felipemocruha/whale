# whale

### Spec

| Symbol | Syntax | Example |
|--------|--------|---------|
| quote | ```(quote form)``` | ```(quote (1 2 3))``` |
| first | ```(first coll)``` | ```(first (quote (1 2 3))``` |
| rest | ```(rest coll)``` | ```(rest (quote (1 2 3))``` |
| cons | ```(cons x seq)``` | ```(cons 4 (quote (1 2 3)))```|
| = | ```(= x y)``` | ```(= "a" "b")``` |
| atom | ```(atom v)``` | ```(atom "abc")``` | 
| cond | ```(cond ...clauses)``` | ```(cond (= 1 2) "ha" (= 1 1) "haha")``` |
| list | ```(list ...items)``` | ```(list 1 "a" (quote (1 2)))``` |
| fn | ```(fn ...args body)``` | ```(fn (x y) (+ x y))``` |
| def | ```(def name val)``` | ```(def x "value")``` |
| defn | ```(defn name ...args body)``` | ```(defn add (x y) (+ x y))``` |


### Descriptions

**quote**: Returns unevaluated form.

**first**: Returns first element from coll.

**rest**: Returns all elements from coll except the first one.

**=**: Returns if x equals y.

**atom**: Returns an atom with value equals v.

**cond**: Evaluates and returns the value associated with the first expression that is yields true.

**list**: Makes a list containing all elements from args.

**fn**: Returns an anonymous function.

**def** Binds a symbol to the current namespace.

**defn**: Defines a function.
