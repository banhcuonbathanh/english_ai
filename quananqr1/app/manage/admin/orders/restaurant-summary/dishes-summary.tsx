import React, { useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import useDeliveryStore from "@/zusstand/delivery/delivery_zustand";
import { toast } from "@/components/ui/use-toast";

const LOG_PATH =
  "quananqr1/app/manage/admin/orders/restaurant-summary/dishes-summary.tsx";

interface OrderDetailedDish {
  dish_id: number;
  quantity: number;
  name: string;
  price: number;
  description: string;
  iamge: string;
  status: string;
}

interface AggregatedDish extends OrderDetailedDish {}

interface NumPadProps {
  dishName: string;
  onSubmit: (value: number) => void;
  onClose: () => void;
}

const NumPad: React.FC<NumPadProps> = ({ dishName, onSubmit, onClose }) => {
  console.log(`[${LOG_PATH}] NumPad rendered for dish:`, dishName);

  const [value, setValue] = useState<string>("");

  const handleNumClick = (num: number) => {
    console.log(`[${LOG_PATH}] NumPad number clicked:`, num);
    setValue((prev) => {
      const newValue = prev + num.toString();
      console.log(`[${LOG_PATH}] NumPad new value:`, newValue);
      return newValue;
    });
  };

  const handleClear = () => {
    console.log(`[${LOG_PATH}] NumPad cleared`);
    setValue("");
  };

  const handleSubmit = () => {
    const parsedValue = parseInt(value || "0", 10);
    console.log(`[${LOG_PATH}] NumPad submitting value:`, parsedValue);
    onSubmit(parsedValue);
    setValue("");
    onClose();
  };

  return (
    <div className="p-4 bg-gray-700">
      <input
        type="text"
        value={value}
        className="w-full p-2 mb-4 text-right text-xl rounded"
        readOnly
      />
      <div className="grid grid-cols-3 gap-2">
        {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((num) => (
          <button
            key={num}
            onClick={() => handleNumClick(num)}
            className="p-4 text-xl rounded"
          >
            {num}
          </button>
        ))}
        <button onClick={handleClear} className="p-4 text-xl rounded">
          C
        </button>
        <button
          onClick={() => handleNumClick(0)}
          className="p-4 text-xl rounded"
        >
          0
        </button>
        <button
          onClick={handleSubmit}
          className="p-4 text-xl rounded bg-blue-500 text-white hover:bg-blue-600"
        >
          ✓
        </button>
      </div>
    </div>
  );
};

const DishSummary: React.FC<{
  dish: AggregatedDish;
  http: any;
  auth: {
    guest: any;
    user: any;
    isGuest: boolean;
  };
  orderStore: {
    tableNumber: number;
    getOrderSummary: () => any;
    clearOrder: () => void;
  };
}> = ({ dish, http, auth, orderStore }) => {
  console.log(`[${LOG_PATH}] DishSummary rendered for dish:`, {
    id: dish.dish_id,
    name: dish.name,
    quantity: dish.quantity
  });

  const [showDetails, setShowDetails] = useState(false);
  const [showNumPad, setShowNumPad] = useState(false);

  const { addDishItem, updateDishQuantity, createDelivery } =
    useDeliveryStore();

  const handleDeliverySubmit = async (quantity: number) => {
    console.log(`[${LOG_PATH}] Handling delivery submit:`, {
      dishId: dish.dish_id,
      dishName: dish.name,
      quantity
    });

    try {
      const deliveryItem = {
        dish_id: dish.dish_id,
        quantity: quantity
      };

      if (dish.quantity === 0) {
        console.log(`[${LOG_PATH}] Adding new dish item:`, deliveryItem);
        addDishItem(deliveryItem);
      } else {
        console.log(`[${LOG_PATH}] Updating dish quantity:`, {
          dishId: dish.dish_id,
          newQuantity: quantity
        });
        updateDishQuantity(dish.dish_id, quantity);
      }

      const deliveryDetails = {
        deliveryAddress: "Default Address",
        deliveryContact: "Default Contact",
        deliveryNotes: "",
        scheduledTime: new Date().toISOString(),
        deliveryFee: 0
      };

      console.log(`[${LOG_PATH}] Creating delivery with details:`, {
        deliveryItem,
        deliveryDetails
      });

      const response = await createDelivery({
        http,
        auth,
        orderStore,
        deliveryDetails
      });

      console.log(`[${LOG_PATH}] Delivery created successfully:`, response);

      toast({
        title: "Success",
        description: `Delivery created for ${quantity} ${dish.name}`
      });

      setShowNumPad(false);
    } catch (error) {
      console.error(`[${LOG_PATH}] Error submitting delivery:`, error);
      toast({
        variant: "destructive",
        title: "Error",
        description:
          error instanceof Error ? error.message : "Failed to create delivery"
      });
    }
  };

  const handleClose = () => {
    console.log(`[${LOG_PATH}] Closing NumPad`);
    setShowNumPad(false);
  };

  const toggleDetails = () => {
    console.log(`[${LOG_PATH}] Toggling dish details:`, {
      dishName: dish.name,
      currentState: showDetails,
      newState: !showDetails
    });
    setShowDetails(!showDetails);
  };

  const handleShowNumPad = () => {
    console.log(`[${LOG_PATH}] Opening NumPad for dish:`, {
      dishName: dish.name,
      currentQuantity: dish.quantity
    });
    setShowNumPad(true);
  };

  return (
    <div className="p-2 mb-2 rounded">
      <div className="flex items-center justify-between">
        <div className="flex-1 cursor-pointer" onClick={toggleDetails}>
          <span className="font-bold">
            {dish.name} : {dish.quantity}
          </span>
        </div>
        <Button
          variant="outline"
          size="sm"
          onClick={handleShowNumPad}
          className="ml-2"
        >
          Delivery
        </Button>
      </div>

      {showDetails && (
        <div className="mt-2 pl-4 text-gray-600">
          <div className="grid grid-cols-2 gap-1">
            <div className="font-medium">Price per Unit:</div>
            <div>${dish.price.toFixed(2)}</div>
            <div className="font-medium">Total Price:</div>
            <div>${(dish.price * dish.quantity).toFixed(2)}</div>
          </div>
        </div>
      )}

      <Dialog open={showNumPad} onOpenChange={setShowNumPad}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Enter Delivery Quantity for {dish.name}</DialogTitle>
          </DialogHeader>
          <NumPad
            dishName={dish.name}
            onSubmit={handleDeliverySubmit}
            onClose={handleClose}
          />
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default DishSummary;
