from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def admin_post01(self):
       self.client.post('/reward',data={'gifcode':'T3WV110C','id':'1'})
    @task
    def admin_post02(self):
       self.client.post('/reward',data={'gifcode':'UZOAK7WS','id':'2'})
    @task
    def admin_post03(self):
       self.client.post('/reward',data={'gifcode':'9GUA959H','id':'3'})