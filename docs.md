Network device mocking

Versions:
0.0.1 Initial version 10/9/2024


Idea:
Generate a docker container that you can ssh into to mimick the behaviour of a network device.
The container runs a single golang program.
The program is an infinite loop of user inputting a command and the program responding with a mocked response.

Use cases:
Device mocking to test network automation
Practise device for network engineers without needing a real device.

The idea:
The whole project is built on the idea to store commands as trees. unique commands will be stored in an array of root nodes. Each node has an output field if it is the final element of a command.  Or an array of possible next values. 
If a value isnt set but rather an argument it will be represented in a argument node. which is largely the same as a command node but instead of a single value it holds the accepted values

Implementation:
Frontend:
Any frontend should be possible. It should translate the mocks to a standardized json format.
the current ideas are:
	json files describing the mock, similar to wiremock
	a gui that allows the user to create trees visualizing the command trees

Translation layer:
The translation layer is the json format the frontend needs to generate for appropriate mocks and responses need to be setup.
A very early implementation of the format:
{
    "commands": [
        {
        "value": "show",
        "output": "no arguments given, use show --help",
        "children": [
            {
            "value": "interface",
            "output": "no arguments given, use show interface --help",
            "children": [
                {
                    "value": "rpd",
                    "outputpath": "/files/showinterfacerpd.txt"
                },
                {
                    "value": "router",
                    "output": "no clue what to put here"
                }
            ]
            }
        ]
        }
    ]
}

Backend:
The backend is the go program that takes in an array of strings, checks if the first values appears in the roots array and then recursively walks down the tree.
When the function gets called with an array of strings of length 1, it outputs the output value of the associated node.
