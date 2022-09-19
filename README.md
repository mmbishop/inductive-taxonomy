# inductive-taxonomy

Code for the article [Of Elephants and Taxonomies](http://curmudgeoncoder.blog/of-elephants-and-taxonomies/) in my Curmudgeon Coder blog. 

This codebase shows how a taxonomy can be built through inductive reasoning using delegation.

See the code in [animal-taxonomy_test.go](https://github.com/mmbishop/inductive-taxonomy/blob/master/src/tests/animal-taxonomy_test.go) for an 
example of how to build a taxonomy through delegation. The objects in the taxonomy each have a prototype. When a request (method call) is sent to an
object, it will respond if it has the local data available to do so. If not, it will delegate the request to its prototype. This will continue 
recursively until a prototype is found that can respond to the request.

Read [Henry Lieberman's seminal paper](https://web.media.mit.edu/~lieber/Lieberary/OOP/Delegation/Delegation.html) on delegation to learn more about it.
