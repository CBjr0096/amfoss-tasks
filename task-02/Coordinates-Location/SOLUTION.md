# Steps I took to solve task-02.

## Making the directory Coordinates-Location.

For making a new directory i used the following command 

>  mkdir Coordinates-Location

## Making the next directory.

First we havd to move to the new directory by using 

>  cd Coordinates-Location

Then we have to make a new directory North by 

>  mkdir North

## Finding DMS values of North.

1.To make the text file NDegree.txt and fill the values by

> echo "25°" > NDegree.txt

2.To make the text file NMinutes.txt and filling the values by 

> echo "5'" > NMinutes.txt

3.To make the text file NSeconds.txt and filling the values by 

> echo "38.1" > NSeconds.txt

then edit the file again to add the quotes(") by 

> nano NSeconds.txt 

4.To make the file NorthCoordinate.txt by
 
> cat * > NorthCoordinate.txt 

5.To copy and rename the file as North.txt by 

> cp NorthCoordinate.txt ~/Documents/Coordinates-Location/North.txt

6. To remove the file created in step 4 from the directory North.

> rm -R NorthCoordinate.txt

## Making the new directory for east and finding the DMS values of East.

1. Change directory to Coordinates-Location from North

> cd ..

2. Making the new directory by 

> mkdir East 

3. Move to the directory and make and fill the files
 EDegree.txt , EMinutes.txt , ESeconds.txt by 

> cd East 

> echo "76°" > EDegree.txt

> echo "29'" > EMinutes.txt

> echo "30.8" > ESeconds.txt

Then to edit Eseconds.txt we use 

> nano ESeconds.txt

4. To make the file EastCoordinate.txt and to copy it to the directory Coordinate-Location and 
remove the original file from the old directory.    

> cat * > Eastcoordinate.txt

> cp EastCoordinate.txt ~/Documents/Coordinate-Location/East.txt

> rm -R EastCoordinate.txt

## Combining files and making it into one single file.

To combine two files i used the command 

> cat North.txt East.txt > Location-Coordinates.txt

## The Screenshot of the google maps 

![](Coordinates-Location/Screenshot.png)

## Making the SOLUTION.md file 

For making the SOLUTION.md file I used 

> cat > SOLUTION.md

> nano SOLUTION.md

## Cloning and pushing to github.

For cloning and pushing to my github repository i used 

> git clone <my repository link>
 
> git add .

> git status
 

> git commit -m <my commit message>

> git log
 
> git push origin main 
 




