import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "@/store/store.ts";
import AdminDashboardLayout from "@/views/admin/AdminDashboardLayout.vue";

const Home = () => import("../views/Home.vue");
const SignUp = () => import("../views/SignUp.vue");
const SignIn = () => import("../views/SignIn.vue");

// Admin dashboard routes
// const AdminDashboardLayout = () => import('../views/admin/AdminDashboardLayout.vue')
const AdminDashboard = () => import("../views/admin/AdminDashboard.vue");
const AdminAllProducts = () => import("../views/admin/AdminAllProducts.vue");
const AdminAddNewProduct = () => import("../views/admin/AddNewProduct.vue");
const AdminOrderList = () => import("../views/admin/AdminOrderList.vue");

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
    meta: { title: "Klock :- Home" },
  },
  {
    path: "/sign-up",
    name: "sign-up",
    component: SignUp,
    meta: { title: "Klock :- Sign Up" },
  },
  {
    path: "/sign-in",
    name: "sign-in",
    component: SignIn,
    meta: { title: "Klock :- Sign In" },
  },
  {
    path: "/dashboard",
    meta: { requiresAuth: true, requiredRole: "ADMIN" },
    component: AdminDashboardLayout,
    children: [
      {
        path: "",
        name: "dashboard",
        component: AdminDashboard,
        meta: { title: "Klock :- Admin Dashboard" },
      },
      {
        path: "all-products",
        name: "all-products",
        component: AdminAllProducts,
        meta: { title: "Klock :- Admin All Products" },
      },
      {
        path: "add-new-product",
        name: "add-new-product",
        component: AdminAddNewProduct,
        meta: {
          title: "Klock :- Admin Add New Product",
          parent: "all-products",
        },
      },
      {
        path: "order-list",
        name: "order-list",
        component: AdminOrderList,
        meta: { title: "Klock :- Admin Order List" },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _, next) => {
  const userStore = useUserStore();
  const accessToken = userStore.accessToken.value;
  const isAuthenticated = () => userStore.accessToken.value !== "";
  const userRole = userStore.userProfile?.Role;
  const requiredRole = to.meta.requiredRole;

  // Protect dashboard route
  if (to.meta.requiresAuth) {
    if (isAuthenticated()) {
      if (userRole === requiredRole) {
        next();
      } else {
        next({ name: "home" });
      }
    } else {
      next({ name: "sign-in" });
    }
  }

  // Prevent user from going to sign-in route if token exists
  else if (accessToken) {
    if ((to.name === "sign-in" || to.name === "sign-up") && isAuthenticated()) {
      if (userRole === "ADMIN") {
        next({ name: "dashboard" });
      } else {
        next({ name: "home" });
      }
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
