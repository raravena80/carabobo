# dinosaurs -- Dinosaurs

Problem is, given two files like this:
```
csv file 1 (Name, length, stance)
---------
Procompsognathus,1.03,bipedal
Triceratops,1.04,quadrupedal
Velociraptor,1.20,bipedal
Brontosaurus,3.0,quadrupedal

csv file 2 (Name, stride, food)
----------
Procompsognathus,2.3,carnivore
Velociraptor,3.21,carnivore
Brontosaurus,0.3,herbivore
Triceratops,1.40,herbivore
```

and speed is: stride * length (for simplicity purposes)

Output only the bipedal dinosaurs sorted by speed

Solutions are in Go:

- dinosaurs.go - Solution in Go. Run with: go run dinosaurs.go
