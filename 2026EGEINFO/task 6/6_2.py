from turtle import *
screensize(2000,2000)
tracer(0)
m=100

left(90)

for i in range(4):
    forward(14*m)
    right(90)
for i in range(5):
    forward(5*m)
    right(45)
penup()

for x in range(-100,100):
    for y in range(-100,100):
        setpos(x*m,y*m)
        dot(5,'red')
done()