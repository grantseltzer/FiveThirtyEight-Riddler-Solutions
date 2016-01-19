from random import randint
import matplotlib.pyplot as plt

# Takes list of previous shots, and returns a new list with the rest of the 98 shots added to it
def takeShot(list):
    listOfShotsTaken = list
    shotsTaken = len(listOfShotsTaken)

    # Recursive base case
    if(shotsTaken < 99):
        mostRecentShot = listOfShotsTaken[-1]
        rando = randint(1,shotsTaken)

        # Determine if shot is made
        if(rando <= mostRecentShot):
            madeShot = True
        else:
            madeShot = False

        # Actions based on outcome of shot
        if(madeShot):
            listOfShotsTaken.append(mostRecentShot + 1)
            takeShot(listOfShotsTaken)
        else:
            listOfShotsTaken.append(mostRecentShot)
            takeShot(listOfShotsTaken)
    else:
        listOfShotsTaken.append(listOfShotsTaken[-1] + 1)
        probOfMakingNextShot = []
        numberShot = 1
        for madeCount in listOfShotsTaken:
            probability = float(madeCount) / float(numberShot)
            probOfMakingNextShot.append(probability)
            numberShot += 1
        shotNumber = range(0,100)
        plt.plot(shotNumber,probOfMakingNextShot)

def main():
    plt.xlabel('Number of Shots Taken')
    plt.ylabel('Probability of Making Next Shot')
    plt.title('The Neurotic Basketball Player\'s Freethrows')
    for i in range(5000):
        takeShot([0,1])
    plt.show()

if __name__ == "__main__":
    main()
