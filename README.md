# k8sutil
This is a command line tool with subcommands for creating and interpreting kubernetes objects.

### TODO
- The utils package is kind of a mess, want to make subdirectories based on each subcommand?
- I eventually want to make the linter extendable. You should be able to progamatically invoke it instead of just as a command-line tool, and you should be able to add your own custom requirements. I like this a lot because since my tool relies onfreehand boolean functions, your tests can literally be whatever you want. You aren't restricted to just set, equal, greaterthan field checks like in kube-lint. This would be really nice. There is a lot more flexibility with this. For example, you could check that a string field belongs to a collection of custom defined strings in your program. Maybe there's not much of a use case for it, but at least the option is there.
    - For this, I would need to finally resign myself to the fact that I will need different rule types for different resource type so I can defer the injection of the relevant resource. (Right now, when I instantiate a rule struct, I am relying on the fact that there is a resource pointer in scope, and this is just not flexible enough. I thought it was a cool idea at first, but I was a little bit wrong)
