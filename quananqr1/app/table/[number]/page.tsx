import { get_dishes } from "@/zusstand/server/dish-controller";
import { DishSelection } from "./component/dish/dishh_list";

import { DishInterface } from "@/schemaValidations/interface/type_dish";
import { SetInterface } from "@/schemaValidations/interface/types_set";
import { get_Sets } from "@/zusstand/server/set-controller";
import { SetCardList } from "./component/set/sets_list";
import OrderSummary from "./component/order/order";
interface TableProps {
  params: { number: string };
  searchParams: { token: string };
}
// This is a server component
export default async function TablePage({ params, searchParams }: TableProps) {
  const number = params.number;
  console.log(" quananqr1/app/test/[number]/page.tsx number", number);
  const token = searchParams.token;

  console.log(" quananqr1/app/test/[number]/page.tsx token", token);
  const dishesData: DishInterface[] = await get_dishes();

  const setsData: SetInterface[] = await get_Sets();
  // const dishes: Dish[] = dishesData;
  // console.log("quananqr1/app/guest/page.tsx dishes.data asdf", setsData);
  return (
    <div className="guest-page">
      <div className="container mx-auto px-4 py-8">
        <img
          src={"/api/placeholder/300/400"}
          className="w-full h-full object-cover rounded-md"
        />
      </div>
      <SetCardList sets={setsData} />

      <DishSelection dishes={dishesData} />

      <OrderSummary number={number} token={token} />
    </div>
  );
}
