from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def admin_post01(self):
       self.client.post('/creategif',data={'gifType':'1','des':'奖励金币','allowTime':'3','valTime':'10m','createName':'aaa','gold':'1000','diamond':'100'})

    @task
    def admin_post02(self):
       self.client.post('/creategif',data={'gifType':'2','des':'奖励金币','allowTime':'2','valTime':'10m','createName':'bbb','gold':'2000','diamond':'500'})

    @task
    def admin_post03(self):
       self.client.post('/creategif',data={'gifType':'3','des':'奖励金币','allowTime':'4','valTime':'10m','createName':'ccc','gold':'3000','diamond':'300'})