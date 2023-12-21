import { ref } from "vue";

// Snackbar
export const success = ref(false);
export const showSnackbar = ref(false);
export const apiResponse = ref("");

// Set colors based on screen size
export const isDesktop = ref(false);
export const textColor = ref("#1D2939");
export const innerColor = ref("black");
export const outerColor = ref("white");

// Show token expiry notification
export const showTokenExpiry = ref(false);

// Button value state
export const buttonValue = ref(false);

// Display mobile dashboard menu
export const displayMenu = ref(false);

// Display user dropdown menu 
export const displayUserDropdownMenu = ref(false);

// Display categories on nav bar
export const displayCategories = ref(false);

// Menu items for header component and mobile menu component
export const menuItems = ref<{ name: string; link?: string; id: string }[]>([
    {
      name: "home",
      link: "/",
      id: "5ba48fa3",
    },
    {
      name: "categories",
      id: "a0a0a0a0",
    },
    {
      name: "Blog",
      link: "/blog",
      id: "4096c3fe",
    },
    {
      name: "about",
      link: "/about",
      id: "9c3b2e28",
    },
    {
      name: "contact",
      link: "/contact",
      id: "a5a5a5a5",
    },
  ]);