# Contributor's Guide
The Automatic Component Toolkit is an open source project.

Contributions are welcome and we are looking for people that can improve existing language bindings or create new bindings or implementation stubs.

You can also contribute by reporting bugs in the [Issue tracker](../../issues), helping review pull requests, participate in discussions about issues and more.

## Filing issues
1. When filing an issue to report errors or problems, make sure to answer these five questions:
	1. Which version of ACT are you using?
		Run <br/>`act.* -v`<br/> to print ACT's version.
	2. Which operating system, programming language(s) and development tools (compiler/interpreter) are you using?
	3. What did you do?
	4. What did you expect to see?
	5. What did you see instead?

2. When contributing to this repository, please first discuss the change you wish to make via issue with the [maintainers](#maintainers) of this repository. This way, we can ensure that there is no overlap between contributions or internal development work.

## Submitting a pull request
When ready to contribute, fork this repository and submit a pull request that references the issue it resolves. Be sure to include a clear and detailed description of the changes you've made so that we can verify them and eventually merge.

ACT follows the [git-flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) branching model. New developments are integrated into the [develop](../../tree/develop)-branch. ACT's maintainers will create releases from the develop-branch when appropriate.

__NOTE__ _Before your code can be accepted into the project you must also sign the Contributor License Agreement (CLA). Please contact the maintainers via automatic-component-toolkit.contributor.agreements@autodesk.com for a copy of the CLA._


## Maintainers
Maintainers are responsible for responding to pull requests and issues, as well as guiding the direction of the project.

We currently have two maintainers:
- Alexander Oster alexander.oster@autodesk.com
- Martin Weismann martin.weismann@autodesk.com

If you've established yourself as an impactful contributor to the project, and are willing take on the extra work, we'd love to have your help maintaining it! Email the current maintainers for details.



Alternatively to 1) build ACT from source ([master](../../tree/master) for a released vesion, [develop](../../tree/develop) for the latest developments):
1. Install go https://golang.org/doc/install
2. Build automaticcomponenttoolkit.go:
<br/>`Build\build.bat` on Windows or <br/>`Build\build.sh` on Unix


## Building ACT from source
1. Install go https://golang.org/doc/install
2. Clone this repo

	```git clone https://github.com/Autodesk/AutomaticComponentToolkit```
3. Check out the develop [develop](../../tree/develop) branch:

	```git checkout develop```

4. Prepare ACT's dependencies:
	1. [lestrrat-go/libxml2](https://github.com/lestrrat-go/libxml2) for XML schema validation:

		1. Install `libxml2` and `libxml2-dev` (e.g. `sudo apt-get install libxml2 libxml2-dev` for Ubuntu, or look at http://xmlsoft.org/index.html for more info)
		2. `go get https://github.com/lestrrat-go/libxml2`

5. Build ACT: 
	- `Build\build.bat` on Windows or
	- `Build\build.sh` on Unix systems
