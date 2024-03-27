import { createRouter, createWebHashHistory } from 'vue-router';
import IndexView from '../views/Index.vue';
import LoginView from '../views/Login.vue';
import CreateView from '../views/Create.vue';
import PostView from '../views/Post.vue';
import InitialReviewView from '../views/InitialReview.vue';
import UserView from '../views/User.vue';
import FinalReviewView from '../views/FinalReview.vue';
import ChartView from '../views/Chart.vue';
import DetailsView from '../views/Details.vue';
import UpdateView from '../views/Update.vue';
import ReviewView from '../views/Review.vue';
import SecondReviewView from '../views/SecondReview.vue';
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView
    },
    {
      path: '/index',
      name: 'index',
      component: IndexView
    },
    {
      path: '/create',
      name: 'create',
      component: CreateView
    },
    {
      path: '/post',
      name: 'post',
      component: PostView
    },
    {
      path: '/initialreview', 
      name: 'initialreview',        
      component: InitialReviewView
    },
    {
      path:'/user',
      name:'user',
      component:UserView
    },
    {
      path: '/chart',
      name: 'chart',
      component: ChartView
    },
    {
      path:'/finalreview',
      name:'finalreview',
      component:FinalReviewView
    },
    {
      path: '/item/details/:id',
      name: 'details',
      component: DetailsView
    },
    {
      path: '/item/update/:id',
      name: 'update',
      component: UpdateView
    },
    {
      path: '/item/review/:id',
      name: 'review',
      component: ReviewView
    },
    {
      path: '/item/review2/:id',
      name: 'review2',
      component: SecondReviewView
    }
  ]
});

export default router;