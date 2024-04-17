import sys

def evaluate_expression(expression):
    try:
        # Evaluate the mathematical expression
        result = eval(expression)
        return result
    except Exception as e:
        return str(e)

if __name__ == "__main__":
    if len(sys.argv) > 1:
        expression = sys.argv[1]
        print(evaluate_expression(expression))
    else:
        print("Please provide a math expression")
