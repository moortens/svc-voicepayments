# svc-voicepayments
We have now build our core product: https://github.com/kimpettersen/svc-payments and we are ready to add product on top of our core service.

First up is svc-voicepayments - a service you can hook up to voice assistants, like Siri, Amazon Alexa and other similar services

This service isn't implemented yet, and that's your job

TODO, in any order you like
 - Implement a method that calls the Remote Server to make a payment (svc-payments)
 - Make it possible to pass command line arguments for the payment information
 - Create a method to retrieve all payments
 - Create a method to confirm a payment
 - Create a method to retrieve a payment by id
 - Create a JSON API to expose this to the external devices
