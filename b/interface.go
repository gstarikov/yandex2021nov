package main

type calcDays = func(X, K int, tasks []int) int

/*
Система устроена следующим образом: для каждой задачи задается дедлайн — день ti.
Если задача не решена до этого момента времени, то в задачу приходит робот и пишет комментарий о том, что задачу надобно закрыть.
Если через X дней задача не решена, то робот приходит снова. Так продолжается до тех пор, пока задача не будет решена.

Алексей каждый день заходит в Yagile и видит сообщения от робота.
Если Алексей не хочет приступать к решению накопленных задач, то он прочитывает все сообщения с помощью одной кнопки и занимается другими делами.
Однако Алексей понимает, что так долго делать нельзя, поэтому он разрешает себе нажимать на эту кнопку ровно K−1 раз,
а на K-й раз садится и решает все задачи разом (даже те, у которых не настал дедлайн).

Определите день, когда Алексей закроет все задачи.
*/
