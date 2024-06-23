# animal
small project to demonstrate module dependencies in go.

module 1. puppy.  is dependant on dog
  it has function
    becute

module 2. dog .   is dependent on animal
  it has functions
    walk & run which uses move from animals.
    bark
    howl

module 3. animal.
  it has functions:
    move
    eat
