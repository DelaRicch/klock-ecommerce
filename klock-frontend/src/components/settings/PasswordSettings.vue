<template>
<form @submit.prevent=handleSubmit class="flex flex-col gap-6">
  <InputField 
  required-tag
  id="current_password"
  label="Current Password"
  type="password"
  placeholder="******"
  :border-color="formData.current_password.error && '#EF0816'"
  @update-value="handleValidateInput($event, 'current_password')"
  :error="formData.current_password.error"
  />
  <InputField 
  required-tag
  id="new-password"
  label="New Password"
  type="password"
  placeholder="******"
  @update-value="handleValidateInput($event, 'password')"
  :error="formData.password.error"
  :border-color="formData.password.error && '#EF0816'"
  />
  <InputField 
  required-tag
  id="confirm-new-password"
  label="Confirm New Password"
  type="password"
  placeholder="******"
  @update-value="handleValidateInput($event, 'confirm_password')"
  :error="formData.confirm_password.error"
  :border-color="formData.confirm_password.error && '#EF0816'"
  />
  <div class="flex self-end">
    <ButtonComponent
    :is-disabled="!isValidForm"
    type="submit" 
    label="Save Changes" 
    color="#FFFFFF" 
    background-color="#000000"
     />
</div>
</form>
</template>
<script setup lang="ts">
import { passwordRegex } from '@/schema/ValidationSchema';
import ButtonComponent from '../ButtonComponent.vue';
import InputField from '../InputField.vue';
import type { formFields } from '@/types';
import { computed, reactive } from 'vue';
import { validateCurrentPassword } from '@/api/user';
import { updatePassword } from '@/api/user';
import { errorApiRequest, successApiRequest } from '@/lib/helperFunctions';

const formData: formFields = reactive({
  current_password: {
    value: "",
    error: "",
  },
  password: {
    value: "",
    error: "",
  },
  confirm_password: {
    value: "",
    error: "",
  },
});

const isValidForm = computed(() => {
  return Object.values(formData).every(
    (field) =>  field.value !== "" && !field.error
  );
});

const handleValidateInput = (value: string, field: string) => {
  if(field === "current_password" && value.length) {
    validateCurrentPassword(value).then(res => {
      if(!res.success) {
        formData.current_password.error = res.message
      }
    }).catch(err => {
      formData.current_password.error = err.message;
    })
  }

  if (field === "password" && !passwordRegex.test(value)) {
    formData[field].error = "Password must be at least 8 characters and include one uppercase letter, one lowercase letter, and one digit";
  } else if (field === "confirm_password") {
    const password = formData["password"].value;
    if (value !== password) {
      formData[field].error = "Passwords do not match";
    } else {
      formData[field].error = "";
    }
  }else {
    formData[field].error = "";
  }
  formData[field].value = value.toString();
};


const handleSubmit = async () => {
  // isSubmitting.value = true;
  const formDataValues: Record<string, string> = {};

  for (const field in formData) {
    // Skip the current_password field
    if (field === 'current_password' || field === 'confirm_password') {
      continue;
    }

    formDataValues[field] = (formData[field] as { value: string }).value;
  }

  // Rest of your code for form submission...
  updatePassword(formDataValues).then(res => {
    successApiRequest(res);
  }).catch(err => {
    errorApiRequest(err);
  })
};

</script>