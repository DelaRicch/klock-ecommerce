
import {createRouter, createWebHistory} from "vue-router";


    const Home = () => import('../views/Home.vue')
 const SignUp = () => import('../views/SignUp.vue')
const SignIn = () => import('../views/SignIn.vue')


  const routes = [
      {path: '/', name: 'home', component: Home, meta: {title: 'Klock :- Home'}},
    {path: '/sign-up', name: 'sign-up', component: SignUp, meta: {title: 'Klock :- Sign Up'}},
      {path: '/sign-in', name: 'sign-in', component: SignIn, meta: {title: 'Klock :- Sign In'}}
  ]

const router = createRouter({
    history: createWebHistory(),
    routes
 })


export default router
