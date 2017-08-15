# in.flux
in.Flux is a little interpreted programming language I created to learn how interpreted languages work

Planned features:
  * Builtin macros:
    ```
    let oneThroughFive = [1, 2, 3, 4, 5]
    let data = {"spam": "eggs", "foo": 5}
    // println prints strings, numbers, and raw versions of objects to stdout, followed by a newline character
    println("Strings and numbers and objects can be printed, followed by a newline \n")
    println(5)
    println(oneThroughFive)
    println(data)
    ```

  * Function chaining:
    ```
    let addStuff = fn(x, y) {
      x + y;
    };
    let numberIsntFive = fn(x) {
      x != 5
    }
    let wasSuccess = fn(res) {
      if res == true {
        println("Success!")
      } else {
        println("Failure :(")
      }
    }

    // Will take the output from addStuff(5, 10) and feed it into numberIsntFive which feeds into wasSuccess
    addStuff(5, 10)->numberIsntFive->wasSuccess
    ```