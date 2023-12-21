export interface formFields {
  [key: string]: {
    value: string;
    error: string;
  };
}
export interface accessTokenType {
  value: string;
  expiry: number;
}

export interface socialAuthType {
  name: string;
  email: string;
  photo: string;
  socialId: string;
  provider: string;
}

export interface UserProfileProps {
  Name: string;
  Email: string;
  UserID: string;
  Role: string;
  Photo: string;
  Phone: string;
}

export interface ProductProps {
  id: number;
  product: string;
  order_id: string;
  amount: string;
  status: string;
  customer_name: string;
  date: string;
}

export interface ErrorObject {
  message: string;
}

interface ProductGalleryImagesProps {
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  ID: number;
  imageUrl: string;
  productId: string;
}

export interface AllProductsProps {
  ProductBrandName: string;
  ProductName: string;
  ProductCategory: string;
  ProductCoverImage: string;
  ProductDescription: string;
  ProductGalleryImages: ProductGalleryImagesProps[];
  ProductID: string;
  ProductPrice: number;
  ProductDiscountPercentage: number;
  ProductQuantity: number;
  ProductsRemaining: number;
  ProductsSold: number;
}
