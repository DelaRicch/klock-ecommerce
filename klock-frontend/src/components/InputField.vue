<template>
  <div
    :class="{ 'max-w-[7rem]': fileClassName, 'w-full': !fileClassName }"
    class="flex flex-col gap-1"
  >
    <div class="flex items-center gap-1">
      <label
        v-if="type !== 'file'"
        :for="id"
        class="line-clamp-1 font-medium text-[#1D2939]"
      >
        {{ label }}
      </label>
      <span v-if="requiredTag" class="font-semibold text-[#EF0816]">*</span>
    </div>
    <div
      v-if="type !== 'file'"
      :class="{
        'h-[2.9rem]': textType === 'input',
        'h-[7rem]': textType === 'textarea',
        className,
      }"
      class="relative"
    >
      <input
        v-if="textType == 'input'"
        @change.prevent="
          $emit('updateValue', ($event.target as HTMLInputElement).value.trim())
        "
        :type="inputType"
        :id="id"
        :step="type === 'number' ? 0.01 : undefined"
        :hidden="type === 'file'"
        :placeholder="placeholder"
        :style="{backgroundColor: backgroundColor, borderColor: borderColor}"
        class="h-full w-full rounded-lg border border-[#667085] px-3 pr-8 outline-none outline-offset-4 placeholder:text-[#667085] hover:border-[#0408E7] hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
      />

      <textarea
        v-if="textType === 'textarea'"
        @input.prevent="
          $emit('updateValue', ($event.target as HTMLInputElement).value.trim())
        "
        :id="id"
        v-model="inputString"
        :placeholder="placeholder"
        :maxlength="maxCharacters"
        class="h-full w-full resize-none rounded-lg border border-[#667085] px-3 pr-8 outline-none outline-offset-4 hover:border-[#0408E7] hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
      ></textarea>
      <div
        v-if="textType === 'textarea'"
        class="flex items-center justify-end pr-1 text-[0.75rem] text-slate-400"
      >
        {{ characterCount }}/{{ maxCharacters }}
      </div>

      <button
        type="button"
        v-if="type === 'password'"
        @click="handleToggleIsHidden"
        class="absolute right-2 top-[50%] translate-y-[-50%] rounded-lg outline-offset-2 focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
      >
        <EyeViewComponent :is-hidden="isHidden" />
      </button>
    </div>
    <div v-if="type === 'file'" class="flex w-full flex-col gap-1">
      <span class="line-clamp-1 font-medium text-[#1D2939]"> {{ label }}</span>
      <input
        v-if="textType == 'input'"
        :type="inputType"
        :id="id"
        hidden
        @change="handleFileUpload"
        accept=".svg, .png, .jpg, .jpeg, .gif"
      />
      <div
        @dragover.prevent
        @dragenter.prevent
        @drop="handleDrop"
        :class="{
          'h-[14rem]': !selectedImage && !fileClassName,
          'h-[5.3rem]': fileClassName,
        }"
        class="relative flex max-h-[15rem] items-center justify-center overflow-clip rounded-lg border border-[#667085] outline-none outline-offset-4 placeholder:text-[#667085] hover:border-[#0408E7] hover:outline-1 hover:outline-[#0408E7] focus:outline-1 focus:outline-[#0408E7] focus:ring-1 focus:ring-[#4B4EFC]"
      >
        <img
          v-if="selectedImage"
          :src="previewImage"
          alt="selected cover image"
          class="h-full w-full object-cover"
          @mouseover="handleDisplayActionButtons(true)"
        />
        <Transition name="fade">
          <div
            v-if="displayActionButtons"
            @mouseleave.stop="handleDisplayActionButtons(false)"
            class="absolute bottom-0 left-0 z-20 flex h-full w-full items-center justify-center gap-12 bg-black/60"
          >
            <label :for="id" class="cursor-pointer">
              <img
                :src="EditIcon"
                alt="Edit Image"
                :class="{ 'w-5': fileClassName, 'w-6 md:w-8': !fileClassName }"
              />
            </label>
            <div @click="handleRemoveImage" class="cursor-pointer">
              <img
                :src="DeleteIcon"
                alt="Delete Image"
                :class="{ 'w-5': fileClassName, 'w-6 md:w-8': !fileClassName }"
              />
            </div>
          </div>
        </Transition>
        <label :for="id">
          <slot v-if="!selectedImage"></slot>
        </label>
      </div>
    </div>
    <span v-if="!!error" class="text-sm text-[#EF0816]">{{ error }}</span>
  </div>
</template>

<script setup lang="ts">
import EyeViewComponent from "./EyeViewComponent.vue";
import DeleteIcon from "@/assets/delete.svg";
import EditIcon from "@/assets/edit.svg";
import { computed, ref, watch } from "vue";
import { errorApiRequest, isValidImage } from "@/lib/helperFunctions";
import { type ErrorObject } from "@/types";

const { label, placeholder, type, id, requiredTag } = defineProps({
  label: {
    type: String,
  },
  placeholder: {
    type: String,
    required: false,
  },
  type: {
    type: String,
    default: "text",
  },
  textType: {
    type: String,
    default: "input",
    validator: (value: string) => ["input", "textarea"].includes(value),
  },
  id: {
    type: String,
    required: true,
  },
  requiredTag: {
    type: Boolean,
    default: false,
  },
  error: {
    type: String,
    default: "",
  },
  maxCharacters: {
    type: Number,
    default: 100,
  },
  className: {
    type: String,
    required: false,
  },
  fileClassName: {
    type: Boolean,
    required: false,
  },
  borderColor: {
    type: String,
  },
  backgroundColor: {
    type: String,
  },
});

const emit = defineEmits(["updateValue"]);

const isHidden = ref(false);
const inputType = ref(type);
const inputString = ref<string>("");
const selectedImage = ref<File | null>(null);
const previewImage = ref("");
const displayActionButtons = ref(false);

const handleToggleIsHidden = () => (isHidden.value = !isHidden.value);

const characterCount = computed(() => inputString.value.trim().length);

const handleDrop = (event: DragEvent) => {
  event.preventDefault();
  const droppedFiles = (event.dataTransfer as DataTransfer).files;
  processFiles(droppedFiles);
};

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const files = target.files as FileList;
  processFiles(files);
};

const processFiles = (files: FileList) => {
  const file = files[0];

  isValidImage(file)
    .then((validatedFile) => {
      selectedImage.value = validatedFile;
      previewImage.value = URL.createObjectURL(validatedFile);
      emit("updateValue", validatedFile);
    })
    .catch((error: ErrorObject) => {
      errorApiRequest(error);
    });
};

const handleRemoveImage = () => {
  selectedImage.value = null;
  previewImage.value = "";
  emit("updateValue", null);
};

const handleDisplayActionButtons = (value: boolean) => {
  displayActionButtons.value = value;
};

watch(isHidden, (value) => {
  inputType.value = value ? "text" : type;
});

watch(selectedImage, (value) => {
  if (!value) {
    displayActionButtons.value = false;
  }
});
</script>

<style lang="css" scoped>
.fade-enter-active,
.fade-leave-active {
  @apply transition-all duration-200 ease-linear;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(100%);
}
</style>
