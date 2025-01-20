Here’s an API documentation for your README file:

---

## API Documentation

### 1. **Create Order**
   - **Endpoint:** `POST /order/create`
   - **Description:** Creates a new order in the system.
   - **Request Body (JSON):**
     ```json
     {
       "User": "user_email",
       "Item": "item_name",
       "Quantity": 2,
       "Price": 100.0
     }
     ```
   - **Response:**
     - **Status:** 201 Created
     - **Body:** 
       ```json
       {
         "ID": "generated_order_id",
         "User": "user_email",
         "Item": "item_name",
         "Quantity": 2,
         "Price": 100.0
       }
       ```
   - **Errors:**
     - 400 Bad Request: If the input JSON is malformed or missing required fields.
     - 500 Internal Server Error: If the order cannot be created.

---

### 2. **Bulk Order Creation**
   - **Endpoint:** `POST /order/create/bulk`
   - **Description:** Creates multiple orders in bulk. The request should include an array of orders.
   - **Request Body (JSON):**
     ```json
     [
       {
         "User": "user1_email",
         "Item": "item1_name",
         "Quantity": 2,
         "Price": 50.0
       },
       {
         "User": "user2_email",
         "Item": "item2_name",
         "Quantity": 1,
         "Price": 150.0
       }
     ]
     ```
   - **Response:**
     - **Status:** 201 Created
     - **Body:**
       ```json
       {
         "ID": "generated_order_id",
         "User": "user1_email",
         "Item": "item1_name",
         "Quantity": 2,
         "Price": 50.0
       }
       ```
   - **Errors:**
     - 400 Bad Request: If the input JSON is malformed or missing required fields.
     - 500 Internal Server Error: If the bulk order creation fails.

---

### 3. **Get Orders**
   - **Endpoint:** `GET /order/get`
   - **Description:** Fetches all orders from the database.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       [
         {
           "ID": "order_id_1",
           "User": "user_email",
           "Item": "item_name",
           "Quantity": 2,
           "Price": 100.0
         },
         {
           "ID": "order_id_2",
           "User": "another_user_email",
           "Item": "another_item_name",
           "Quantity": 1,
           "Price": 200.0
         }
       ]
       ```
   - **Errors:**
     - 500 Internal Server Error: If fetching orders fails.

---

### 4. **Update Order**
   - **Endpoint:** `PUT /order/update`
   - **Description:** Updates an existing order in the system.
   - **Request Body (JSON):**
     ```json
     {
       "ID": "order_id",
       "User": "updated_user_email",
       "Item": "updated_item_name",
       "Quantity": 3,
       "Price": 120.0
     }
     ```
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "ID": "updated_order_id",
         "User": "updated_user_email",
         "Item": "updated_item_name",
         "Quantity": 3,
         "Price": 120.0
       }
       ```
   - **Errors:**
     - 400 Bad Request: If the input JSON is malformed.
     - 404 Not Found: If the order to be updated does not exist.
     - 500 Internal Server Error: If the update operation fails.

---

### 5. **Delete Order**
   - **Endpoint:** `DELETE /order/delete`
   - **Description:** Deletes an existing order by its ID.
   - **Query Parameter:**
     - `id`: The unique identifier of the order to delete.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "message": "Order deleted successfully"
       }
       ```
   - **Errors:**
     - 400 Bad Request: If the ID is not a valid number.
     - 404 Not Found: If the order with the given ID does not exist.
     - 500 Internal Server Error: If the deletion fails.

---

### Common Error Responses:
   - **400 Bad Request:** Invalid input or missing required parameters.
   - **404 Not Found:** The resource (order) was not found.
   - **500 Internal Server Error:** An unexpected error occurred during the operation.

--- 

This documentation covers all the essential endpoints and expected responses. Ensure your request bodies are properly formatted to avoid errors.

![{8719470B-644B-444D-BD92-3A71DAD6F4E8}](https://github.com/user-attachments/assets/23b3fc16-fdaa-4c08-927f-ea08e67bcc4d)
![{00457677-E3A5-4337-B2EB-ED3095FB6F21}](https://github.com/user-attachments/assets/c97fc0ab-12bd-4ba3-a07d-57a9f28e007a)
![{ED7098EC-7EA7-4EDE-A88C-7581EE668B75}](https://github.com/user-attachments/assets/57126fb9-f5c2-4fdf-8752-a922df303059)
![{741FBCE7-708F-4C65-8DAF-12FEF857AAA9}](https://github.com/user-attachments/assets/fa51372d-e6c0-4453-a45b-d347392b05e8)

