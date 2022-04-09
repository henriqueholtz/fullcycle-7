import { EntityNotFoundError } from 'typeorm/error/EntityNotFoundError';
import { ArgumentsHost, Catch, ExceptionFilter } from '@nestjs/common';
import { Response } from 'express';

@Catch(EntityNotFoundError)
export class ModelNotFoundExceptionFilter implements ExceptionFilter {
  catch(exception: EntityNotFoundError, host: ArgumentsHost) {
    const context = host.switchToHttp();
    const response = context.getResponse<Response>();

    return response.status(404).json({
      error: {
        error: '404 - Not Found',
        message: exception.message,
      },
    });
  }
}
