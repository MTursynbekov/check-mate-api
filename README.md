# Check-Mate

App is created for tracking activities with your friends, family members, coworkers and other people.

As database Postgres was used.

This app provides functionality to: <br>
1. Add contacts.
1. Store interactions(converstions) with them
1. Store their birthdays and add reminders about them

<br>
Backend is written in Golang. <br>

## Routes

- POST /signup 
- POST /signin
<br>
- POST /contacts 
- GET /contacts/:id
<br>
- POST /chat
- GET /chat/:id
- GET /users/:userId/contacts Get all contacts of user
- GET /users/:userId/contacts/:contactId Get certain contact of the user