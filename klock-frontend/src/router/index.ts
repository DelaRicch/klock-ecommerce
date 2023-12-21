import { createRouter, createWebHistory } from "vue-router";
import AdminDashboardLayout from "@/views/admin/AdminDashboardLayout.vue";
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useUserStore } from "@/stores/user";

const Home = () => import("@/views/Home.vue");
const SignUp = () => import("@/views/SignUp.vue");
const SignIn = () => import("@/views/SignIn.vue");

// user routes
const AccountSettings = () => import("@/views/AccountSettings.vue");

// Admin dashboard routes
// const AdminDashboardLayout = () => import('@/views/admin/AdminDashboardLayout.vue')
const AdminDashboard = () => import("@/views/admin/AdminDashboard.vue");
const AdminAllProducts = () => import("@/views/admin/AdminAllProducts.vue");
const AdminAddNewProduct = () => import("@/views/admin/AddNewProduct.vue");
const AdminOrderList = () => import("@/views/admin/AdminOrderList.vue");

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

  // User route
  {
    path: "/account-settings",
    name: "account-settings",
    component: AccountSettings,
    meta: {title: "Klock :- Account Settings", requiresAuth: true}
  },

  // Admin route
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
  const {accessToken, userProfile} = storeToRefs(useUserStore());
  const isAuthenticated = computed(() => !!accessToken.value.value);
  const userRole = userProfile?.value?.Role;
  const requiredRole = to.meta.requiredRole;

  // Protect dashboard route
  if (to.meta.requiresAuth) {
    if (isAuthenticated.value) {
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
    if ((to.name === "sign-in" || to.name === "sign-up") && isAuthenticated.value) {
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
