import numpy as np
import matplotlib.pyplot as pyplot


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

EXPECTED_VALUE = np.pi

if __name__ == "__main__":
    
    k = 50


    b_pts = [[] for x in range(100)]
    u_pts = [[] for x in range(100)]
    c_pts = [[] for x in range(100)]
    d_pts = [[] for x in range(100)]

    for j in range(5000):
        line = input()
        text = line.strip().split()
        b_pts[j % 100].append(Point(1000*(j % 100 + 1), float(text[0])))
        u_pts[j % 100].append(Point(1000*(j % 100 + 1), float(text[1])))
        c_pts[j % 100].append(Point(1000*(j % 100 + 1), float(text[2])))
        d_pts[j % 100].append(Point(1000*(j % 100 + 1), float(text[3])))

    pyplot.clf()
    pyplot.grid(False)

    # fig, axs = pyplot.subplots()

    # axs.set_yticks(np.arange(0, 2, 0.2))


    plot_pts = []
    for i in range(100):
        n = (i + 1)*1000
        b = [point.y for point in b_pts[i]]
        avgb = sum(b)/len(b)

        plot_pts.append(Point(n, avgb/np.sqrt(n)))

    x = [point.x for point in plot_pts]
    y = [point.y for point in plot_pts]

    pyplot.plot(x, y)

    pyplot.show()
    

    

