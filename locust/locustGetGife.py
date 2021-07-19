from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def admin_post01(self):
       self.client.post("/getGifcode",data={'gifcode':'MJWJHYRV'})

    @task
    def admin_post02(self):
       self.client.post("/getGifcode",data={'gifcode':'FX34FG1P'})

    @task
    def admin_post03(self):
       self.client.post("/getGifcode",data={'gifcode':'GSUMLBUN'})