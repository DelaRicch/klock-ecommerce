
<template>
  <div class="w-max overflow-hidden relative whitespace-nowrap" ref="starsContainer" :style="{
    pointerEvents: readonly ? 'none' : 'auto',
  }">
    <div class="rating-outer absolute top-0 left-0">
      <component :is="StarIcon" v-for="i in numberOfStars" :key="i" class="rating-icon mx-0.5" />
    </div>
    <div class="rating-inner">
      <component :is="StarIcon" v-for="i in numberOfStars" :key="i" class="rating-icon mx-0.5" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watchEffect } from "vue";
import StarIcon from "@/assets/StarIcon.vue";

interface Emit {
  (e: "update:modelValue", v: number): void;
}

const props = defineProps({
  numberOfStars: {
    type: Number,
    default: 5,
  },
  starSize: {
    type: Number,
    default: 24,
  },
  modelValue: {
    type: Number,
    default: 0,
  },
  starColor: {
    type: String,
    default: "#000000",
  },
  inactiveColor: {
    type: String,
    default: "#FFFFFF",
  },
  readonly: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits<Emit>();

const utils = {
  rounded(value: number, decimalPlaces: number) {
    const power = Math.pow(10, decimalPlaces);
    return Math.round(value * power) / power;
  },
};

const starsContainer = ref<HTMLDivElement>();

const rating = computed({
  get() {
    return props.modelValue;
  },
  set(newVal) {
    const roundedVal = utils.rounded(newVal, 1);
    newVal = Math.ceil(roundedVal * 2) / 2;

    emit("update:modelValue", newVal);
  },
});


function handleSelectRating (this: HTMLDivElement, e: MouseEvent) {
  if (props.readonly) return;
  const rect = this.getBoundingClientRect();
  const { pageX } = e;
  const relativeX = pageX - rect.left;
  const offsetWidth = rect.width;

  const numberOfStars = props.numberOfStars;

  const result = (relativeX / offsetWidth) * numberOfStars;

  rating.value = result;
}

const percent = computed(() => {
  const normalizedRating =
    rating.value < 0
      ? 0
      : rating.value > props.numberOfStars
        ? props.numberOfStars
        : rating.value;

  return (normalizedRating / props.numberOfStars) * 100;
});

watchEffect(() => {
  const styleValues = {
    "--innerColor": props.inactiveColor,
    "--outerColor": props.starColor,
    "--outerWidth": `${percent.value}%`,
    "--iconSize": `${props.starSize}px`,
  };

  for (const [key, value] of Object.entries(styleValues)) {
    starsContainer.value?.style.setProperty(key, value);
  }
});

onMounted(() => {
  starsContainer.value?.addEventListener("click", handleSelectRating );
});

onBeforeUnmount(() => {
  starsContainer.value?.removeEventListener("click", handleSelectRating );
});
</script>

<style scoped>

.rating-outer {
  width: var(--outerWidth);
  max-width: 100%;
  overflow: hidden;
  color: var(--outerColor);
  transition: width 320ms cubic-bezier(0.075, 0.82, 0.165, 1);
}

.rating-inner {
  color: var(--innerColor);
}

.rating-icon {
  fill: currentColor;
  width: var(--iconSize);
  aspect-ratio: 1;
  cursor: pointer;
  display: inline-block;
}
</style>