import type { UseFormReturn } from "react-hook-form";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Input } from "./ui/input";

interface fieldProps {
  name: string;
  type: string;
  label: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  form: UseFormReturn<any, any, any>;
}
export default function FormInput(props: fieldProps) {
  return (
    <FormField
      control={props.form.control}
      name={props.name}
      render={({ field }) => (
        <FormItem className="px-3">
          <FormLabel>{props.label}</FormLabel>
          <FormControl>
            <Input
              className="dark:bg-secondary bg-secondary"
              autoComplete="off"
              type={props.type}
              placeholder={props.name}
              {...field}
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
