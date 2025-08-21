import random
from locust import HttpUser, task, between


class RetailerUser(HttpUser):
    """
    User class that simulates a user browsing a retailer's product API.
    """
    # wait_time defines how long a simulated user will wait between executing tasks.
    # Here, it's a random time between 1 and 3 seconds. This is crucial to
    # simulate realistic user behavior and avoid overwhelming the server instantly.
    wait_time = between(1, 3)

    def on_start(self):
        """
        on_start is called when a user is started.
        We can use this to set up any initial data for the user.
        """
        # Let's create a list of potential product IDs this user might query.
        self.product_ids = [
            "apple-iphone-15", "samsung-galaxy-s24", "google-pixel-8", "sony-wh-1000xm5"]

    # This task will be picked 3 times more often than the post_product task.
    @task(3)
    def get_product(self):
        """
        Simulates a user fetching information for a random product.
        This represents a read operation (GET request).
        """
        # Pick a random product ID from the list
        product_id = random.choice(self.product_ids)

        # Make a GET request to the /products endpoint with the chosen product_id
        # The 'name' parameter groups all requests to this endpoint under a single
        # entry in the Locust UI, making the statistics cleaner.
        self.client.get(f"/products/{product_id}", name="/products/[id]")

    @task(1)  # This task has a weight of 1, making it less frequent.
    def post_product(self):
        """
        Simulates a user or system adding a new product.
        This represents a write operation (POST request).
        """
        # Define the new product data we want to send.
        new_product = {
            "id": "new-product-" + str(random.randint(1000, 9999)),
            "name": "Super Gadget",
            "price": 999.99
        }

        # Make a POST request to the /products endpoint, sending the new_product
        # data as a JSON payload.
        self.client.post("/products", json=new_product)
