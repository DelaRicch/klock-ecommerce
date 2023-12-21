
<template>
    <form @submit.prevent="handleSubmit" class="flex flex-col gap-6">
        <InputField 
        id="name"
        label="Name"
        type="text"
        :placeholder="userProfile?.Name"
        @update-value="handleValidateInput($event, 'name')"
        />
        <InputField 
        id="email"
        label="Email"
        type="email"
        :placeholder="userProfile?.Email"
        :error="formData.email.error"
        @update-value="handleValidateInput($event, 'email')"
        />
        <InputField 
        id="phone"
        label="Contact"
        type="tel"
        placeholder="+233 240 590 212"
        @update-value="handleValidateInput($event, 'phone')"
        />
        <InputField 
        id="location"
        label="Location"
        type="text"
        placeholder="St. Augustine FL 32003, United States"
        @update-value="handleValidateInput($event, 'location')"
        />
        <div class="flex self-end">
            <ButtonComponent 
            type="submit" 
            label="Save Changes" 
            color="#FFFFFF" 
            background-color="#000000"
            :is-disabled="isInvalidForm"
             />
        </div>
    </form>
</template>
<script setup lang="ts">
import { useUserStore } from '@/stores/user';
import InputField from '../InputField.vue';
import ButtonComponent from '../ButtonComponent.vue';
import { computed, reactive } from 'vue';
import { emailRegex } from '@/schema/ValidationSchema';
import type { formFields } from '@/types';
import { updateUser } from '@/api/user';
import { errorApiRequest, handleUserProfile, successApiRequest } from '@/lib/helperFunctions';

const {userProfile} = useUserStore();

const formData: formFields = reactive({
  name: {
    value: "",
    error: "",
  },
  email: {
    value: "",
    error: "",
  },
  phone: {
    value: "",
    error: "",
  },
  location: {
    value: "",
    error: "",
  },
});

const isInvalidForm = computed(() => {
  for (const field in formData) {
    if (formData[field].error) {
      return true;
    }
  }
  return false;
});

const handleValidateInput = (value: string, field: string) => {
  if (field === "email" && !emailRegex.test(value)) {
    formData[field].error = "Invalid email format";
  } else {
    formData[field].error = "";
  }

  formData[field].value = value.toString();
};

const handleSubmit = async () => {
  const formDataValues: Record<string, string> = {};
  for (const field in formData) {
    formDataValues[field] = (formData[field] as { value: string }).value;
  }

  updateUser(formDataValues).then((res) => {
    console.log(res);
    handleUserProfile(res.user);

    let result = {
      success: res.success,
      message: res.message
    }
    successApiRequest(result);
  }).catch((err) => {
    console.log(err);
    errorApiRequest(err);
  });
}

</script>