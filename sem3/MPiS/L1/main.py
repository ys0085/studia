import numpy as np
import matplotlib.pyplot as pyplot


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

EXPECTED_VALUE = np.pi

if __name__ == "__main__":
    pts = []
    avgs = []
    line = input()
    x = 50
    while line != "":
        text = line.strip().split()
        values = []
        for t in text:
            v = float(t)
            values.append(v)
            pts.append(Point(x, v))
        avg = sum(values)/len(values)
        avgs.append(Point(x, avg))
        x += 50
        line = input()

    pyplot.clf()
    pyplot.grid(True)

    x = [point.x for point in pts]
    y = [point.y for point in pts]
    
    pyplot.scatter(x, y, color="blue", s=1.5)

    avgx = [point.x for point in avgs]
    avgy = [point.y for point in avgs]
    
    pyplot.scatter(avgx, avgy, color="red", s=8)

    pyplot.axhline(y=EXPECTED_VALUE, color="green")

    pyplot.show()
    
    

    

